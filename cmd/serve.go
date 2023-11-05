/*
Copyright © 2023 saltfishpr <526191197@qq.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"time"

	"learning/internal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start timezone http server",

	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		i := do.New()

		config := internal.NewConfig()
		if err := viper.Unmarshal(config); err != nil {
			return err
		}
		do.ProvideValue[*internal.Config](i, config)

		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			AddSource: true,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.SourceKey {
					s := a.Value.Any().(*slog.Source)
					s.File = path.Join(path.Base(path.Dir(s.File)), path.Base(s.File))
				}
				return a
			},
		}))
		do.ProvideValue[*slog.Logger](i, logger)

		do.Provide[internal.Repo](i, func(i *do.Injector) (internal.Repo, error) {
			return internal.NewMysqlRepo(i)
		})
		do.Provide[*internal.Handler](i, internal.NewHandler)

		e := echo.New()
		e.HideBanner = true

		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{}))
		e.Use(internal.Logger(logger))
		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 5 * time.Second,
		}))

		e.GET("/liveness", func(c echo.Context) error {
			return c.String(http.StatusOK, "OK")
		})
		handler := do.MustInvoke[*internal.Handler](i)
		e.POST("/timezones", handler.CreateTimeZone)
		e.PATCH("/timezones/:id", handler.UpdateTimeZone)
		e.DELETE("/timezones/:id", handler.DeleteTimeZone)
		e.GET("/timezones/:id", handler.GetTimeZone)
		e.GET("/timezones", handler.ListTimeZones)
		e.GET("/timezones:load", handler.LoadTimeZones)

		return e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
