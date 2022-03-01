package graph
import (
    "strings"
    "errors"
)
// 
type ResultantAppStateDetail int

const (
    PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL ResultantAppStateDetail = iota
    MINIMUMDISKSPACENOTMET_RESULTANTAPPSTATEDETAIL
    MINIMUMOSVERSIONNOTMET_RESULTANTAPPSTATEDETAIL
    MINIMUMPHYSICALMEMORYNOTMET_RESULTANTAPPSTATEDETAIL
    MINIMUMLOGICALPROCESSORCOUNTNOTMET_RESULTANTAPPSTATEDETAIL
    MINIMUMCPUSPEEDNOTMET_RESULTANTAPPSTATEDETAIL
    PLATFORMNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    FILESYSTEMREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    REGISTRYREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    POWERSHELLSCRIPTREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
    SUPERSEDINGAPPSNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    NOADDITIONALDETAILS_RESULTANTAPPSTATEDETAIL
    DEPENDENCYFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
    DEPENDENCYWITHREQUIREMENTSNOTMET_RESULTANTAPPSTATEDETAIL
    DEPENDENCYPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    DEPENDENCYWITHAUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
    SUPERSEDEDAPPUNINSTALLFAILED_RESULTANTAPPSTATEDETAIL
    SUPERSEDEDAPPUNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    REMOVINGSUPERSEDEDAPPS_RESULTANTAPPSTATEDETAIL
    IOSAPPSTOREUPDATEFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
    VPPAPPHASUPDATEAVAILABLE_RESULTANTAPPSTATEDETAIL
    USERREJECTEDUPDATE_RESULTANTAPPSTATEDETAIL
    UNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    SUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    SUPERSEDEDAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    SEEINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
    AUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
    MANAGEDAPPNOLONGERPRESENT_RESULTANTAPPSTATEDETAIL
    USERREJECTEDINSTALL_RESULTANTAPPSTATEDETAIL
    USERISNOTLOGGEDINTOAPPSTORE_RESULTANTAPPSTATEDETAIL
    UNTARGETEDSUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
    APPREMOVEDBYSUPERSEDENCE_RESULTANTAPPSTATEDETAIL
    SEEUNINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
    PENDINGREBOOT_RESULTANTAPPSTATEDETAIL
    INSTALLINGDEPENDENCIES_RESULTANTAPPSTATEDETAIL
    CONTENTDOWNLOADED_RESULTANTAPPSTATEDETAIL
)

func (i ResultantAppStateDetail) String() string {
    return []string{"PROCESSORARCHITECTURENOTAPPLICABLE", "MINIMUMDISKSPACENOTMET", "MINIMUMOSVERSIONNOTMET", "MINIMUMPHYSICALMEMORYNOTMET", "MINIMUMLOGICALPROCESSORCOUNTNOTMET", "MINIMUMCPUSPEEDNOTMET", "PLATFORMNOTAPPLICABLE", "FILESYSTEMREQUIREMENTNOTMET", "REGISTRYREQUIREMENTNOTMET", "POWERSHELLSCRIPTREQUIREMENTNOTMET", "SUPERSEDINGAPPSNOTAPPLICABLE", "NOADDITIONALDETAILS", "DEPENDENCYFAILEDTOINSTALL", "DEPENDENCYWITHREQUIREMENTSNOTMET", "DEPENDENCYPENDINGREBOOT", "DEPENDENCYWITHAUTOINSTALLDISABLED", "SUPERSEDEDAPPUNINSTALLFAILED", "SUPERSEDEDAPPUNINSTALLPENDINGREBOOT", "REMOVINGSUPERSEDEDAPPS", "IOSAPPSTOREUPDATEFAILEDTOINSTALL", "VPPAPPHASUPDATEAVAILABLE", "USERREJECTEDUPDATE", "UNINSTALLPENDINGREBOOT", "SUPERSEDINGAPPSDETECTED", "SUPERSEDEDAPPSDETECTED", "SEEINSTALLERRORCODE", "AUTOINSTALLDISABLED", "MANAGEDAPPNOLONGERPRESENT", "USERREJECTEDINSTALL", "USERISNOTLOGGEDINTOAPPSTORE", "UNTARGETEDSUPERSEDINGAPPSDETECTED", "APPREMOVEDBYSUPERSEDENCE", "SEEUNINSTALLERRORCODE", "PENDINGREBOOT", "INSTALLINGDEPENDENCIES", "CONTENTDOWNLOADED"}[i]
}
func ParseResultantAppStateDetail(v string) (interface{}, error) {
    result := PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
    switch strings.ToUpper(v) {
        case "PROCESSORARCHITECTURENOTAPPLICABLE":
            result = PROCESSORARCHITECTURENOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "MINIMUMDISKSPACENOTMET":
            result = MINIMUMDISKSPACENOTMET_RESULTANTAPPSTATEDETAIL
        case "MINIMUMOSVERSIONNOTMET":
            result = MINIMUMOSVERSIONNOTMET_RESULTANTAPPSTATEDETAIL
        case "MINIMUMPHYSICALMEMORYNOTMET":
            result = MINIMUMPHYSICALMEMORYNOTMET_RESULTANTAPPSTATEDETAIL
        case "MINIMUMLOGICALPROCESSORCOUNTNOTMET":
            result = MINIMUMLOGICALPROCESSORCOUNTNOTMET_RESULTANTAPPSTATEDETAIL
        case "MINIMUMCPUSPEEDNOTMET":
            result = MINIMUMCPUSPEEDNOTMET_RESULTANTAPPSTATEDETAIL
        case "PLATFORMNOTAPPLICABLE":
            result = PLATFORMNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "FILESYSTEMREQUIREMENTNOTMET":
            result = FILESYSTEMREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "REGISTRYREQUIREMENTNOTMET":
            result = REGISTRYREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "POWERSHELLSCRIPTREQUIREMENTNOTMET":
            result = POWERSHELLSCRIPTREQUIREMENTNOTMET_RESULTANTAPPSTATEDETAIL
        case "SUPERSEDINGAPPSNOTAPPLICABLE":
            result = SUPERSEDINGAPPSNOTAPPLICABLE_RESULTANTAPPSTATEDETAIL
        case "NOADDITIONALDETAILS":
            result = NOADDITIONALDETAILS_RESULTANTAPPSTATEDETAIL
        case "DEPENDENCYFAILEDTOINSTALL":
            result = DEPENDENCYFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
        case "DEPENDENCYWITHREQUIREMENTSNOTMET":
            result = DEPENDENCYWITHREQUIREMENTSNOTMET_RESULTANTAPPSTATEDETAIL
        case "DEPENDENCYPENDINGREBOOT":
            result = DEPENDENCYPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "DEPENDENCYWITHAUTOINSTALLDISABLED":
            result = DEPENDENCYWITHAUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
        case "SUPERSEDEDAPPUNINSTALLFAILED":
            result = SUPERSEDEDAPPUNINSTALLFAILED_RESULTANTAPPSTATEDETAIL
        case "SUPERSEDEDAPPUNINSTALLPENDINGREBOOT":
            result = SUPERSEDEDAPPUNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "REMOVINGSUPERSEDEDAPPS":
            result = REMOVINGSUPERSEDEDAPPS_RESULTANTAPPSTATEDETAIL
        case "IOSAPPSTOREUPDATEFAILEDTOINSTALL":
            result = IOSAPPSTOREUPDATEFAILEDTOINSTALL_RESULTANTAPPSTATEDETAIL
        case "VPPAPPHASUPDATEAVAILABLE":
            result = VPPAPPHASUPDATEAVAILABLE_RESULTANTAPPSTATEDETAIL
        case "USERREJECTEDUPDATE":
            result = USERREJECTEDUPDATE_RESULTANTAPPSTATEDETAIL
        case "UNINSTALLPENDINGREBOOT":
            result = UNINSTALLPENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "SUPERSEDINGAPPSDETECTED":
            result = SUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "SUPERSEDEDAPPSDETECTED":
            result = SUPERSEDEDAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "SEEINSTALLERRORCODE":
            result = SEEINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
        case "AUTOINSTALLDISABLED":
            result = AUTOINSTALLDISABLED_RESULTANTAPPSTATEDETAIL
        case "MANAGEDAPPNOLONGERPRESENT":
            result = MANAGEDAPPNOLONGERPRESENT_RESULTANTAPPSTATEDETAIL
        case "USERREJECTEDINSTALL":
            result = USERREJECTEDINSTALL_RESULTANTAPPSTATEDETAIL
        case "USERISNOTLOGGEDINTOAPPSTORE":
            result = USERISNOTLOGGEDINTOAPPSTORE_RESULTANTAPPSTATEDETAIL
        case "UNTARGETEDSUPERSEDINGAPPSDETECTED":
            result = UNTARGETEDSUPERSEDINGAPPSDETECTED_RESULTANTAPPSTATEDETAIL
        case "APPREMOVEDBYSUPERSEDENCE":
            result = APPREMOVEDBYSUPERSEDENCE_RESULTANTAPPSTATEDETAIL
        case "SEEUNINSTALLERRORCODE":
            result = SEEUNINSTALLERRORCODE_RESULTANTAPPSTATEDETAIL
        case "PENDINGREBOOT":
            result = PENDINGREBOOT_RESULTANTAPPSTATEDETAIL
        case "INSTALLINGDEPENDENCIES":
            result = INSTALLINGDEPENDENCIES_RESULTANTAPPSTATEDETAIL
        case "CONTENTDOWNLOADED":
            result = CONTENTDOWNLOADED_RESULTANTAPPSTATEDETAIL
        default:
            return 0, errors.New("Unknown ResultantAppStateDetail value: " + v)
    }
    return &result, nil
}
func SerializeResultantAppStateDetail(values []ResultantAppStateDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
