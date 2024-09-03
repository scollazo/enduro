// Code generated by goa v3.15.2, DO NOT EDIT.
//
// enduro HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	package_c "github.com/artefactual-sdps/enduro/internal/api/gen/http/package_/client"
	storagec "github.com/artefactual-sdps/enduro/internal/api/gen/http/storage/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `package (monitor-request|monitor|list|show|preservation-actions|confirm|reject|move|move-status|upload)
storage (create|submit|update|download|move|move-status|reject|show|locations|add-location|show-location|location-packages)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` package monitor-request --token "abc123"` + "\n" +
		os.Args[0] + ` storage create --body '{
      "aip_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "location_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "name": "abc123",
      "object_key": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "status": "in_review"
   }' --token "abc123"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	package_Configurer *package_c.ConnConfigurer,
) (goa.Endpoint, any, error) {
	var (
		package_Flags = flag.NewFlagSet("package", flag.ContinueOnError)

		package_MonitorRequestFlags     = flag.NewFlagSet("monitor-request", flag.ExitOnError)
		package_MonitorRequestTokenFlag = package_MonitorRequestFlags.String("token", "", "")

		package_MonitorFlags      = flag.NewFlagSet("monitor", flag.ExitOnError)
		package_MonitorTicketFlag = package_MonitorFlags.String("ticket", "", "")

		package_ListFlags                   = flag.NewFlagSet("list", flag.ExitOnError)
		package_ListNameFlag                = package_ListFlags.String("name", "", "")
		package_ListAipIDFlag               = package_ListFlags.String("aip-id", "", "")
		package_ListEarliestCreatedTimeFlag = package_ListFlags.String("earliest-created-time", "", "")
		package_ListLatestCreatedTimeFlag   = package_ListFlags.String("latest-created-time", "", "")
		package_ListLocationIDFlag          = package_ListFlags.String("location-id", "", "")
		package_ListStatusFlag              = package_ListFlags.String("status", "", "")
		package_ListCursorFlag              = package_ListFlags.String("cursor", "", "")
		package_ListTokenFlag               = package_ListFlags.String("token", "", "")

		package_ShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		package_ShowIDFlag    = package_ShowFlags.String("id", "REQUIRED", "Identifier of package to show")
		package_ShowTokenFlag = package_ShowFlags.String("token", "", "")

		package_PreservationActionsFlags     = flag.NewFlagSet("preservation-actions", flag.ExitOnError)
		package_PreservationActionsIDFlag    = package_PreservationActionsFlags.String("id", "REQUIRED", "Identifier of package to look up")
		package_PreservationActionsTokenFlag = package_PreservationActionsFlags.String("token", "", "")

		package_ConfirmFlags     = flag.NewFlagSet("confirm", flag.ExitOnError)
		package_ConfirmBodyFlag  = package_ConfirmFlags.String("body", "REQUIRED", "")
		package_ConfirmIDFlag    = package_ConfirmFlags.String("id", "REQUIRED", "Identifier of package to look up")
		package_ConfirmTokenFlag = package_ConfirmFlags.String("token", "", "")

		package_RejectFlags     = flag.NewFlagSet("reject", flag.ExitOnError)
		package_RejectIDFlag    = package_RejectFlags.String("id", "REQUIRED", "Identifier of package to look up")
		package_RejectTokenFlag = package_RejectFlags.String("token", "", "")

		package_MoveFlags     = flag.NewFlagSet("move", flag.ExitOnError)
		package_MoveBodyFlag  = package_MoveFlags.String("body", "REQUIRED", "")
		package_MoveIDFlag    = package_MoveFlags.String("id", "REQUIRED", "Identifier of package to move")
		package_MoveTokenFlag = package_MoveFlags.String("token", "", "")

		package_MoveStatusFlags     = flag.NewFlagSet("move-status", flag.ExitOnError)
		package_MoveStatusIDFlag    = package_MoveStatusFlags.String("id", "REQUIRED", "Identifier of package to move")
		package_MoveStatusTokenFlag = package_MoveStatusFlags.String("token", "", "")

		package_UploadFlags           = flag.NewFlagSet("upload", flag.ExitOnError)
		package_UploadContentTypeFlag = package_UploadFlags.String("content-type", "multipart/form-data; boundary=goa", "")
		package_UploadTokenFlag       = package_UploadFlags.String("token", "", "")
		package_UploadStreamFlag      = package_UploadFlags.String("stream", "REQUIRED", "path to file containing the streamed request body")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageCreateFlags     = flag.NewFlagSet("create", flag.ExitOnError)
		storageCreateBodyFlag  = storageCreateFlags.String("body", "REQUIRED", "")
		storageCreateTokenFlag = storageCreateFlags.String("token", "", "")

		storageSubmitFlags     = flag.NewFlagSet("submit", flag.ExitOnError)
		storageSubmitBodyFlag  = storageSubmitFlags.String("body", "REQUIRED", "")
		storageSubmitAipIDFlag = storageSubmitFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageSubmitTokenFlag = storageSubmitFlags.String("token", "", "")

		storageUpdateFlags     = flag.NewFlagSet("update", flag.ExitOnError)
		storageUpdateAipIDFlag = storageUpdateFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageUpdateTokenFlag = storageUpdateFlags.String("token", "", "")

		storageDownloadFlags     = flag.NewFlagSet("download", flag.ExitOnError)
		storageDownloadAipIDFlag = storageDownloadFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageDownloadTokenFlag = storageDownloadFlags.String("token", "", "")

		storageMoveFlags     = flag.NewFlagSet("move", flag.ExitOnError)
		storageMoveBodyFlag  = storageMoveFlags.String("body", "REQUIRED", "")
		storageMoveAipIDFlag = storageMoveFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageMoveTokenFlag = storageMoveFlags.String("token", "", "")

		storageMoveStatusFlags     = flag.NewFlagSet("move-status", flag.ExitOnError)
		storageMoveStatusAipIDFlag = storageMoveStatusFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageMoveStatusTokenFlag = storageMoveStatusFlags.String("token", "", "")

		storageRejectFlags     = flag.NewFlagSet("reject", flag.ExitOnError)
		storageRejectAipIDFlag = storageRejectFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageRejectTokenFlag = storageRejectFlags.String("token", "", "")

		storageShowFlags     = flag.NewFlagSet("show", flag.ExitOnError)
		storageShowAipIDFlag = storageShowFlags.String("aip-id", "REQUIRED", "Identifier of AIP")
		storageShowTokenFlag = storageShowFlags.String("token", "", "")

		storageLocationsFlags     = flag.NewFlagSet("locations", flag.ExitOnError)
		storageLocationsTokenFlag = storageLocationsFlags.String("token", "", "")

		storageAddLocationFlags     = flag.NewFlagSet("add-location", flag.ExitOnError)
		storageAddLocationBodyFlag  = storageAddLocationFlags.String("body", "REQUIRED", "")
		storageAddLocationTokenFlag = storageAddLocationFlags.String("token", "", "")

		storageShowLocationFlags     = flag.NewFlagSet("show-location", flag.ExitOnError)
		storageShowLocationUUIDFlag  = storageShowLocationFlags.String("uuid", "REQUIRED", "Identifier of location")
		storageShowLocationTokenFlag = storageShowLocationFlags.String("token", "", "")

		storageLocationPackagesFlags     = flag.NewFlagSet("location-packages", flag.ExitOnError)
		storageLocationPackagesUUIDFlag  = storageLocationPackagesFlags.String("uuid", "REQUIRED", "Identifier of location")
		storageLocationPackagesTokenFlag = storageLocationPackagesFlags.String("token", "", "")
	)
	package_Flags.Usage = package_Usage
	package_MonitorRequestFlags.Usage = package_MonitorRequestUsage
	package_MonitorFlags.Usage = package_MonitorUsage
	package_ListFlags.Usage = package_ListUsage
	package_ShowFlags.Usage = package_ShowUsage
	package_PreservationActionsFlags.Usage = package_PreservationActionsUsage
	package_ConfirmFlags.Usage = package_ConfirmUsage
	package_RejectFlags.Usage = package_RejectUsage
	package_MoveFlags.Usage = package_MoveUsage
	package_MoveStatusFlags.Usage = package_MoveStatusUsage
	package_UploadFlags.Usage = package_UploadUsage

	storageFlags.Usage = storageUsage
	storageCreateFlags.Usage = storageCreateUsage
	storageSubmitFlags.Usage = storageSubmitUsage
	storageUpdateFlags.Usage = storageUpdateUsage
	storageDownloadFlags.Usage = storageDownloadUsage
	storageMoveFlags.Usage = storageMoveUsage
	storageMoveStatusFlags.Usage = storageMoveStatusUsage
	storageRejectFlags.Usage = storageRejectUsage
	storageShowFlags.Usage = storageShowUsage
	storageLocationsFlags.Usage = storageLocationsUsage
	storageAddLocationFlags.Usage = storageAddLocationUsage
	storageShowLocationFlags.Usage = storageShowLocationUsage
	storageLocationPackagesFlags.Usage = storageLocationPackagesUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "package":
			svcf = package_Flags
		case "storage":
			svcf = storageFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "package":
			switch epn {
			case "monitor-request":
				epf = package_MonitorRequestFlags

			case "monitor":
				epf = package_MonitorFlags

			case "list":
				epf = package_ListFlags

			case "show":
				epf = package_ShowFlags

			case "preservation-actions":
				epf = package_PreservationActionsFlags

			case "confirm":
				epf = package_ConfirmFlags

			case "reject":
				epf = package_RejectFlags

			case "move":
				epf = package_MoveFlags

			case "move-status":
				epf = package_MoveStatusFlags

			case "upload":
				epf = package_UploadFlags

			}

		case "storage":
			switch epn {
			case "create":
				epf = storageCreateFlags

			case "submit":
				epf = storageSubmitFlags

			case "update":
				epf = storageUpdateFlags

			case "download":
				epf = storageDownloadFlags

			case "move":
				epf = storageMoveFlags

			case "move-status":
				epf = storageMoveStatusFlags

			case "reject":
				epf = storageRejectFlags

			case "show":
				epf = storageShowFlags

			case "locations":
				epf = storageLocationsFlags

			case "add-location":
				epf = storageAddLocationFlags

			case "show-location":
				epf = storageShowLocationFlags

			case "location-packages":
				epf = storageLocationPackagesFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "package":
			c := package_c.NewClient(scheme, host, doer, enc, dec, restore, dialer, package_Configurer)
			switch epn {
			case "monitor-request":
				endpoint = c.MonitorRequest()
				data, err = package_c.BuildMonitorRequestPayload(*package_MonitorRequestTokenFlag)
			case "monitor":
				endpoint = c.Monitor()
				data, err = package_c.BuildMonitorPayload(*package_MonitorTicketFlag)
			case "list":
				endpoint = c.List()
				data, err = package_c.BuildListPayload(*package_ListNameFlag, *package_ListAipIDFlag, *package_ListEarliestCreatedTimeFlag, *package_ListLatestCreatedTimeFlag, *package_ListLocationIDFlag, *package_ListStatusFlag, *package_ListCursorFlag, *package_ListTokenFlag)
			case "show":
				endpoint = c.Show()
				data, err = package_c.BuildShowPayload(*package_ShowIDFlag, *package_ShowTokenFlag)
			case "preservation-actions":
				endpoint = c.PreservationActions()
				data, err = package_c.BuildPreservationActionsPayload(*package_PreservationActionsIDFlag, *package_PreservationActionsTokenFlag)
			case "confirm":
				endpoint = c.Confirm()
				data, err = package_c.BuildConfirmPayload(*package_ConfirmBodyFlag, *package_ConfirmIDFlag, *package_ConfirmTokenFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = package_c.BuildRejectPayload(*package_RejectIDFlag, *package_RejectTokenFlag)
			case "move":
				endpoint = c.Move()
				data, err = package_c.BuildMovePayload(*package_MoveBodyFlag, *package_MoveIDFlag, *package_MoveTokenFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = package_c.BuildMoveStatusPayload(*package_MoveStatusIDFlag, *package_MoveStatusTokenFlag)
			case "upload":
				endpoint = c.Upload()
				data, err = package_c.BuildUploadPayload(*package_UploadContentTypeFlag, *package_UploadTokenFlag)
				if err == nil {
					data, err = package_c.BuildUploadStreamPayload(data, *package_UploadStreamFlag)
				}
			}
		case "storage":
			c := storagec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = storagec.BuildCreatePayload(*storageCreateBodyFlag, *storageCreateTokenFlag)
			case "submit":
				endpoint = c.Submit()
				data, err = storagec.BuildSubmitPayload(*storageSubmitBodyFlag, *storageSubmitAipIDFlag, *storageSubmitTokenFlag)
			case "update":
				endpoint = c.Update()
				data, err = storagec.BuildUpdatePayload(*storageUpdateAipIDFlag, *storageUpdateTokenFlag)
			case "download":
				endpoint = c.Download()
				data, err = storagec.BuildDownloadPayload(*storageDownloadAipIDFlag, *storageDownloadTokenFlag)
			case "move":
				endpoint = c.Move()
				data, err = storagec.BuildMovePayload(*storageMoveBodyFlag, *storageMoveAipIDFlag, *storageMoveTokenFlag)
			case "move-status":
				endpoint = c.MoveStatus()
				data, err = storagec.BuildMoveStatusPayload(*storageMoveStatusAipIDFlag, *storageMoveStatusTokenFlag)
			case "reject":
				endpoint = c.Reject()
				data, err = storagec.BuildRejectPayload(*storageRejectAipIDFlag, *storageRejectTokenFlag)
			case "show":
				endpoint = c.Show()
				data, err = storagec.BuildShowPayload(*storageShowAipIDFlag, *storageShowTokenFlag)
			case "locations":
				endpoint = c.Locations()
				data, err = storagec.BuildLocationsPayload(*storageLocationsTokenFlag)
			case "add-location":
				endpoint = c.AddLocation()
				data, err = storagec.BuildAddLocationPayload(*storageAddLocationBodyFlag, *storageAddLocationTokenFlag)
			case "show-location":
				endpoint = c.ShowLocation()
				data, err = storagec.BuildShowLocationPayload(*storageShowLocationUUIDFlag, *storageShowLocationTokenFlag)
			case "location-packages":
				endpoint = c.LocationPackages()
				data, err = storagec.BuildLocationPackagesPayload(*storageLocationPackagesUUIDFlag, *storageLocationPackagesTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// packageUsage displays the usage of the package command and its subcommands.
func package_Usage() {
	fmt.Fprintf(os.Stderr, `The package service manages packages being transferred to a3m.
Usage:
    %[1]s [globalflags] package COMMAND [flags]

COMMAND:
    monitor-request: Request access to the /monitor WebSocket
    monitor: Obtain access to the /monitor WebSocket
    list: List all stored packages
    show: Show package by ID
    preservation-actions: List all preservation actions by ID
    confirm: Signal the package has been reviewed and accepted
    reject: Signal the package has been reviewed and rejected
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package
    upload: Upload a package to trigger an ingest workflow

Additional help:
    %[1]s package COMMAND --help
`, os.Args[0])
}
func package_MonitorRequestUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package monitor-request -token STRING

Request access to the /monitor WebSocket
    -token STRING: 

Example:
    %[1]s package monitor-request --token "abc123"
`, os.Args[0])
}

func package_MonitorUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package monitor -ticket STRING

Obtain access to the /monitor WebSocket
    -ticket STRING: 

Example:
    %[1]s package monitor --ticket "abc123"
`, os.Args[0])
}

func package_ListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package list -name STRING -aip-id STRING -earliest-created-time STRING -latest-created-time STRING -location-id STRING -status STRING -cursor STRING -token STRING

List all stored packages
    -name STRING: 
    -aip-id STRING: 
    -earliest-created-time STRING: 
    -latest-created-time STRING: 
    -location-id STRING: 
    -status STRING: 
    -cursor STRING: 
    -token STRING: 

Example:
    %[1]s package list --name "abc123" --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --earliest-created-time "1970-01-01T00:00:01Z" --latest-created-time "1970-01-01T00:00:01Z" --location-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --status "in progress" --cursor "abc123" --token "abc123"
`, os.Args[0])
}

func package_ShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package show -id UINT -token STRING

Show package by ID
    -id UINT: Identifier of package to show
    -token STRING: 

Example:
    %[1]s package show --id 1 --token "abc123"
`, os.Args[0])
}

func package_PreservationActionsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package preservation-actions -id UINT -token STRING

List all preservation actions by ID
    -id UINT: Identifier of package to look up
    -token STRING: 

Example:
    %[1]s package preservation-actions --id 1 --token "abc123"
`, os.Args[0])
}

func package_ConfirmUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package confirm -body JSON -id UINT -token STRING

Signal the package has been reviewed and accepted
    -body JSON: 
    -id UINT: Identifier of package to look up
    -token STRING: 

Example:
    %[1]s package confirm --body '{
      "location_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5"
   }' --id 1 --token "abc123"
`, os.Args[0])
}

func package_RejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package reject -id UINT -token STRING

Signal the package has been reviewed and rejected
    -id UINT: Identifier of package to look up
    -token STRING: 

Example:
    %[1]s package reject --id 1 --token "abc123"
`, os.Args[0])
}

func package_MoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move -body JSON -id UINT -token STRING

Move a package to a permanent storage location
    -body JSON: 
    -id UINT: Identifier of package to move
    -token STRING: 

Example:
    %[1]s package move --body '{
      "location_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5"
   }' --id 1 --token "abc123"
`, os.Args[0])
}

func package_MoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package move-status -id UINT -token STRING

Retrieve the status of a permanent storage location move of the package
    -id UINT: Identifier of package to move
    -token STRING: 

Example:
    %[1]s package move-status --id 1 --token "abc123"
`, os.Args[0])
}

func package_UploadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] package upload -content-type STRING -token STRING -stream STRING

Upload a package to trigger an ingest workflow
    -content-type STRING: 
    -token STRING: 
    -stream STRING: path to file containing the streamed request body

Example:
    %[1]s package upload --content-type "multipart/form-data; boundary=goa" --token "abc123" --stream "goa.png"
`, os.Args[0])
}

// storageUsage displays the usage of the storage command and its subcommands.
func storageUsage() {
	fmt.Fprintf(os.Stderr, `The storage service manages the storage of packages.
Usage:
    %[1]s [globalflags] storage COMMAND [flags]

COMMAND:
    create: Create a new package
    submit: Start the submission of a package
    update: Signal that a package submission is complete
    download: Download package by AIPID
    move: Move a package to a permanent storage location
    move-status: Retrieve the status of a permanent storage location move of the package
    reject: Reject a package
    show: Show package by AIPID
    locations: List locations
    add-location: Create a storage location
    show-location: Show location by UUID
    location-packages: List all the packages stored in the location with UUID

Additional help:
    %[1]s storage COMMAND --help
`, os.Args[0])
}
func storageCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage create -body JSON -token STRING

Create a new package
    -body JSON: 
    -token STRING: 

Example:
    %[1]s storage create --body '{
      "aip_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "location_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "name": "abc123",
      "object_key": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5",
      "status": "in_review"
   }' --token "abc123"
`, os.Args[0])
}

func storageSubmitUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage submit -body JSON -aip-id STRING -token STRING

Start the submission of a package
    -body JSON: 
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage submit --body '{
      "name": "abc123"
   }' --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage update -aip-id STRING -token STRING

Signal that a package submission is complete
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage update --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageDownloadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage download -aip-id STRING -token STRING

Download package by AIPID
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage download --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageMoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move -body JSON -aip-id STRING -token STRING

Move a package to a permanent storage location
    -body JSON: 
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage move --body '{
      "location_id": "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5"
   }' --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageMoveStatusUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage move-status -aip-id STRING -token STRING

Retrieve the status of a permanent storage location move of the package
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage move-status --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageRejectUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage reject -aip-id STRING -token STRING

Reject a package
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage reject --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show -aip-id STRING -token STRING

Show package by AIPID
    -aip-id STRING: Identifier of AIP
    -token STRING: 

Example:
    %[1]s storage show --aip-id "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageLocationsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage locations -token STRING

List locations
    -token STRING: 

Example:
    %[1]s storage locations --token "abc123"
`, os.Args[0])
}

func storageAddLocationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage add-location -body JSON -token STRING

Create a storage location
    -body JSON: 
    -token STRING: 

Example:
    %[1]s storage add-location --body '{
      "config": {
         "Type": "s3",
         "Value": "{\"bucket\":\"abc123\",\"endpoint\":\"abc123\",\"key\":\"abc123\",\"path_style\":false,\"profile\":\"abc123\",\"region\":\"abc123\",\"secret\":\"abc123\",\"token\":\"abc123\"}"
      },
      "description": "abc123",
      "name": "abc123",
      "purpose": "aip_store",
      "source": "minio"
   }' --token "abc123"
`, os.Args[0])
}

func storageShowLocationUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show-location -uuid STRING -token STRING

Show location by UUID
    -uuid STRING: Identifier of location
    -token STRING: 

Example:
    %[1]s storage show-location --uuid "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}

func storageLocationPackagesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage location-packages -uuid STRING -token STRING

List all the packages stored in the location with UUID
    -uuid STRING: Identifier of location
    -token STRING: 

Example:
    %[1]s storage location-packages --uuid "d1845cb6-a5ea-474a-9ab8-26f9bcd919f5" --token "abc123"
`, os.Args[0])
}
