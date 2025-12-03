package main

import (
	"fmt"
	"io"
	"log"

	"github.com/pkg/errors"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

type SteamApp struct {
	AppID     string
	StoreName string
}

func getSteamRunningApps() ([]*SteamApp, error) {
	// 遍历 HKEY_CURRENT_USER\Software\Valve\Steam\Apps
	appsKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Valve\Steam\Apps`, registry.READ)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer appsKey.Close()
	appIDs, err := appsKey.ReadSubKeyNames(65536)
	if err != nil && err != io.EOF {
		return nil, errors.WithStack(err)
	}

	var runningApps []*SteamApp
	for _, appIDStr := range appIDs {
		app, err := getAppIfRunning(appsKey, appIDStr)
		if err != nil {
			log.Printf("Error checking app %s: %v", appIDStr, err)
			continue
		}
		if app != nil {
			runningApps = append(runningApps, app)
		}
	}
	return runningApps, nil
}

func getAppIfRunning(appsKey registry.Key, appIDStr string) (*SteamApp, error) {
	appKey, err := registry.OpenKey(appsKey, appIDStr, registry.READ)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer appKey.Close()
	v, _, err := appKey.GetIntegerValue("Running")
	if err != nil {
		if err == windows.ERROR_FILE_NOT_FOUND {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	if v != 1 {
		return nil, nil
	}
	storeName, _, err := appKey.GetStringValue("Name")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &SteamApp{
		AppID:     appIDStr,
		StoreName: storeName,
	}, nil
}

func main() {
	apps, err := getSteamRunningApps()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	for _, app := range apps {
		fmt.Printf("Running AppID: %s, Name: %s\n", app.AppID, app.StoreName)
	}
}
