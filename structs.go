package eapi

import (
	"errors"
)

type TranslationConfig struct {
	Supported   []string
	Definitions map[string][]string
}

type TranslationChange struct {
	Code      string
	Formatted string
}

type Translation struct {
    prefix   string
    markdown TranslationMDData

	Language        string
	UseCustom       bool
	CustomResponses Custom
}

type TranslationMDData struct {
    Start string
    End   string
}

type Pages struct {
	Pages []string

	length int
}

type Custom struct {
	Role_Exist           string `json:"ErrRoleExist"`
	Role_Permission      string `json:"ErrRolePermission"`
	Master_User_Add      string `json:"AddMaster"`
	Master_User_Remove   string `json:"DelMaster"`
	Prefix_Change        string `json:"Prefix"`
	Greet_Enable         string `json:"GreetEnable"`
	Greet_Disable        string `json:"GreetDisable"`
	Greet_Message_Change string `json:"GreetMessage"`
	Bye_Enable           string `json:"ByeEnable"`
	Bye_Disable          string `json:"ByeDisable"`
	Bye_Message_Change   string `json:"ByeMessage"`
	AutoRole_Enable      string `json:"AutoRoleEnable"`
	AutoRole_Disable     string `json:"AutoRoleDisable"`
	AutoRole_Role_Change string `json:"AutoRole"`
	Role_Give            string `json:"Give"`
	Role_Take            string `json:"Take"`
	AllowLinks           string `json:"AllowLinks"`
	DenyLinks            string `json:"DenyLinks"`
	Antilink_Warn        string `json:"AntilinkWarn"`
	AntiLink_Kick        string `json:"AntiLinkKick"`
	AnitLink_Ban         string `json:"AnitLinkBan"`
	Kick                 string `json:"Kick"`
	Kick_Reason          string `json:"KickReason"`
	Ban_Add              string `json:"BanAdd"`
	Ban_Add_Reason       string `json:"BanAddReason"`
}

type CategorySubHelp struct {
	SubHelp struct {
		Help  string `json:"Help"`
		Files string `json:"Files"`
	} `json:"SubHelp"`
}

type CategoryErrors struct {
	Errors struct {
		Process_Run             string `json:"Process_Run"`
		Process_Regex           string `json:"Process_Regex"`
		Invalid_State           string `json:"Invalid_State"`
		Command_Disabled        string `json:"Command_Disabled"`
		Command_NotExist        string `json:"Command_NotExist"`
		Guild_WriteFile         string `json:"Guild_WriteFile"`
		Guild_State             string `json:"Guild_State"`
		Role_Exist              string `json:"Role_Exist"`
		Role_Permission         string `json:"Role_Permission"`
		Channel_Exist           string `json:"Channel_Exist"`
		Channel_Provide         string `json:"Channel_Provide"`
		Channel_Not_Provided    string `json:"Channel_Not_Provided"`
		User_Exist              string `json:"User_Exist"`
		User_IsNotMaster        string `json:"User_IsNotMaster"`
		User_IsNotOwner         string `json:"User_IsNotOwner"`
		Message_NoMention       string `json:"Message_NoMention"`
		Message_Permission_Send string `json:"Message_Permission_Send"`
		Message_Permission_Get  string `json:"Message_Permission_Get"`
		Message_Pin             string `json:"Message_Pin"`
	} `json:"Errors"`
}

type CategoryAttachments struct {
	Attachments struct {
		Database_Install     string `json:"Database_Install"`
		Database_Update      string `json:"Database_Update"`
		AutoResponse_Install string `json:"AutoResponse_Install"`
		Snippet_Install      string `json:"Snippet_Install"`
	} `json:"Attachments"`
}

type CategoryGeneral struct {
	General struct {
		Contact                           string `json:"Contact"`
		Contact_DM                        string `json:"Contact_DM"`
		Lock_Question                     string `json:"Lock_Question"`
		Lock_Confirmation                 string `json:"Lock_Confirmation"`
		Unlock                            string `json:"Unlock"`
		Help                              string `json:"Help"`
		Enable_All                        string `json:"Enable_All"`
		Enable_Command                    string `json:"Enable_Command"`
		Enable_CommandGroup               string `json:"Enable_CommandGroup"`
		Disable_Command                   string `json:"Disable_Command"`
		Disable_CommandGroup              string `json:"Disable_CommandGroup"`
		Language_Error                    string `json:"Language_Error"`
		Language_Invalid                  string `json:"Language_Invalid"`
		Language_NotSupported             string `json:"Language_NotSupported"`
		Language_Change                   string `json:"Language_Change"`
		Prefix_Error                      string `json:"Prefix_Error"`
		Prefix_Change                     string `json:"Prefix_Change"`
		Prefix_TooLong                    string `json:"Prefix_TooLong"`
		Prefix_Reserved                   string `json:"Prefix_Reserved"`
		Master_Explanation                string `json:"Master_Explanation"`
		Master_Mode_Set                   string `json:"Master_Mode_Set"`
		Master_Mode_Explanation           string `json:"Master_Mode_Explanation"`
		Master_Role_Change                string `json:"Master_Role_Change"`
		Master_Role_Error                 string `json:"Master_Role_Error"`
		Master_Permission_Change          string `json:"Master_Permission_Change"`
		Master_Permission_Change_Multiple string `json:"Master_Permission_Change_Multiple"`
		Master_Permission_Error           string `json:"Master_Permission_Error"`
		Master_Permission_TooMuch_Error   string `json:"Master_Permission_TooMuch_Error"`
		Master_User_Add                   string `json:"Master_User_Add"`
		Master_User_Add_Multiple          string `json:"Master_User_Add_Multiple"`
		Master_User_Remove                string `json:"Master_User_Remove"`
		Master_User_Remove_Multiple       string `json:"Master_User_Remove_Multiple"`
		Master_User_Error                 string `json:"Master_User_Error"`
		Master_User_WrongMode_Error       string `json:"Master_User_WrongMode_Error"`
	} `json:"General"`
}

type CategoryEvents struct {
	Events struct {
		Greet_Enable         string `json:"Greet_Enable"`
		Greet_Disable        string `json:"Greet_Disable"`
		Greet_Message_Change string `json:"Greet_Message_Change"`
		Greet_Channel_Change string `json:"Greet_Channel_Change"`
		Greet_Channel_DM     string `json:"Greet_Channel_DM"`
		Bye_Enable           string `json:"Bye_Enable"`
		Bye_Disable          string `json:"Bye_Disable"`
		Bye_Message_Change   string `json:"Bye_Message_Change"`
		Bye_Channel_Change   string `json:"Bye_Channel_Change"`
		Bye_Channel_DM       string `json:"Bye_Channel_DM"`
	} `json:"Events"`
}

type CategoryFilters struct {
	Filters struct {
		AllowLinks                   string `json:"AllowLinks"`
		DenyLinks                    string `json:"DenyLinks"`
		AntiLink_Warn                string `json:"AntiLink_Warn"`
		AntiLink_Kick                string `json:"AntiLink_Kick"`
		AntiLink_Ban                 string `json:"AntiLink_Ban"`
		WordFilter_Enable            string `json:"WordFilter_Enable"`
		WordFilter_Disable           string `json:"WordFilter_Disable"`
		WordFilter_Error             string `json:"WordFilter_Error"`
		WordFilter_Error_Sub         string `json:"WordFilter_Error_Sub"`
		WordFilter_Add               string `json:"WordFilter_Add"`
		WordFilter_Delete            string `json:"WordFilter_Delete"`
		WordFilter_Empty             string `json:"WordFilter_Empty"`
		WordFilter_Dupe              string `json:"WordFilter_Dupe"`
		WordFilter_Warn              string `json:"WordFilter_Warn"`
		WordFilter_Kick              string `json:"WordFilter_Kick"`
		WordFilter_Ban               string `json:"WordFilter_Ban"`
		NameFilter_Enable            string `json:"NameFilter_Enable"`
		NameFilter_Disable           string `json:"NameFilter_Disable"`
		NameFilter_Error             string `json:"NameFilter_Error"`
		NameFilter_Error_Sub         string `json:"NameFilter_Error_Sub"`
		NameFilter_Error_Sub2        string `json:"NameFilter_Error_Sub2"`
		NameFilter_Error_Sub3        string `json:"NameFilter_Error_Sub3"`
		NameFilter_Add               string `json:"NameFilter_Add"`
		NameFilter_Delete            string `json:"NameFilter_Delete"`
		NameFilter_Empty             string `json:"NameFilter_Empty"`
		NameFilter_Kick              string `json:"NameFilter_Kick"`
		FilterAction_Error           string `json:"FilterAction_Error"`
		FilterAction_Mode_Error      string `json:"FilterAction_Mode_Error"`
		FilterAction_Mode_NotAllowed string `json:"FilterAction_Mode_NotAllowed"`
		FilterAction_Mode_Change     string `json:"FilterAction_Mode_Change"`
	} `json:"Filters"`
}

type CategoryActions struct {
	Actions struct {
		AntiLink_Kick   string `json:"AntiLink_Kick"`
		AntiLink_Ban    string `json:"AntiLink_Ban"`
		WordFilter_Kick string `json:"WordFilter_Kick"`
		WordFilter_Ban  string `json:"WordFilter_Ban"`
		NameFilter_Kick string `json:"NameFilter_Kick"`
		Kick_NoReason   string `json:"Kick_NoReason"`
		Ban_NoReason    string `json:"Ban_NoReason"`
	} `json:"Actions"`
}

type CategoryCommandsDiscord struct {
	CommandsDiscord struct {
		Kick                          string `json:"Kick"`
		Kick_Reason                   string `json:"Kick_Reason"`
		Kick_Error                    string `json:"Kick_Error"`
		Ban_Add                       string `json:"Ban_Add"`
		Ban_Add_Reason                string `json:"Ban_Add_Reason"`
		Ban_Remove                    string `json:"Ban_Remove"`
		Ban_Add_Error                 string `json:"Ban_Add_Error"`
		Ban_Remove_Error              string `json:"Ban_Remove_Error"`
		SafeRole_Error                string `json:"SafeRole_Error"`
		SafeRole_Enable               string `json:"SafeRole_Enable"`
		SafeRole_Disable              string `json:"SafeRole_Disable"`
		Role_Give                     string `json:"Role_Give"`
		Role_Give_SafeRole_Position   string `json:"Role_Give_SafeRole_Position"`
		Role_Give_SafeRole_Manage     string `json:"Role_Give_SafeRole_Manage"`
		Role_Give_Error_Syntax        string `json:"Role_Give_Error_Syntax"`
		Role_Give_Error_Echo_Position string `json:"Role_Give_Error_Echo_Position"`
		Role_Give_Error               string `json:"Role_Give_Error"`
		Role_Take                     string `json:"Role_Take"`
		Role_Take_SafeRole_Position   string `json:"Role_Take_SafeRole_Position"`
		Role_Take_SafeRole_Manage     string `json:"Role_Take_SafeRole_Manage"`
		Role_Take_Error_Syntax        string `json:"Role_Take_Error_Syntax"`
		Role_Take_Error_Echo_Position string `json:"Role_Take_Error_Echo_Position"`
		Role_Take_Error               string `json:"Role_Take_Error"`
		Clear_Error_Syntax            string `json:"Clear_Error_Syntax"`
		Clear_Error_Amount            string `json:"Clear_Error_Amount"`
	} `json:"Commands_Discord"`
}

type CategoryCommandsEcho struct {
	CommandsEcho struct {
		ErrorReport_Enable         string `json:"ErrorReport_Enable"`
		ErrorReport_Disable        string `json:"ErrorReport_Disable"`
		ErrorReport_Channel_Change string `json:"ErrorReport_Channel_Change"`
		ErrorReport_Error          string `json:"ErrorReport_Error"`
		AutoRole_Enable            string `json:"AutoRole_Enable"`
		AutoRole_Disable           string `json:"AutoRole_Disable"`
		AutoRole_Role_Change       string `json:"AutoRole_Role_Change"`
		AutoRole_Channel_Change    string `json:"AutoRole_Channel_Change"`
		AutoRole_Event             string `json:"AutoRole_Event"`
		AutoNick_Enable            string `json:"AutoNick_Enable"`
		AutoNick_Disable           string `json:"AutoNick_Disable"`
		AutoNick_Change            string `json:"AutoNick_Change"`
		AutoNick_Error_Syntax      string `json:"AutoNick_Error_Syntax"`
		Nickname                   string `json:"Nickname"`
		Nickname_Error_Syntax      string `json:"Nickname_Error_Syntax"`
		Nickname_Error             string `json:"Nickname_Error"`
	} `json:"Commands_Echo"`
}

type CategoryCommandsMisc struct {
	CommandsMisc struct {
		Command_ID          string   `json:"Command_ID"`
		Command_ID_Error    string   `json:"Command_ID_Error"`
		LocateIP            string   `json:"LocateIP"`
		LocateIP_Error      string   `json:"LocateIP_Error"`
		Track_Invites       string   `json:"Track_Invites"`
		Track_Invites_Error string   `json:"Track_Invites_Error"`
		EightBall           []string `json:"EightBall"`
		Ask_For_Prefix      string   `json:"Ask_For_Prefix"`
	} `json:"Commands_Misc"`
}

type CategoryProfiles struct {
	Profiles struct {
		Info                 string `json:"Info"`
		List                 string `json:"List"`
		NotExist_Error       string `json:"NotExist_Error"`
		IsExist_Error        string `json:"IsExist_Error"`
        Create               string `json:"Create"`
        Create_Syntax_Error  string `json:"Create_Syntax_Error"`
        Delete               string `json:"Delete"`
        Delete_Syntax_Error  string `json:"Delete_Syntax_Error"`
		Delete_Default_Error string `json:"Delete_Default_Error"`
        Import               string `json:"Import"`
        Import_Syntax_Error  string `json:"Import_Syntax_Error"`
		Import_Same_Error    string `json:"Import_Same_Error"`
		Check_Locked         string `json:"Check_Locked"`
		Check_NotLocked      string `json:"Check_NotLocked"`
        Lock                 string `json:"Lock"`
        Lock_Syntax_Error    string `json:"Lock_Syntax_Error"`
		Lock_Default_Error   string `json:"Lock_Default_Error"`
		Unlock               string `json:"Unlock"`
	} `json:"Profiles"`
}

// Error constants.
var (
	ErrLanguageNotSupported = errors.New("Provided language is not supported.")
	ErrTranslationExist     = errors.New("Requested translation does not exist.")
	ErrTranslationEmpty     = errors.New("No translation present for this language.")
	/*ErrTranslationReadFile                 = errors.New("Could not read file.")
	  ErrTranslationTemplateJson             = errors.New("Unable to load json into Template.")
	  ErrTranslationTemplateEmpty            = errors.New("No Template data in interface 'Template'.")
	  ErrTranslationDataEmpty                = errors.New("No language data in interface 'Data'.")
	  ErrTranslationLanguageUndefined        = errors.New("No language defined.")
	  ErrTranslationTemplateNoCategories     = errors.New("Template has no categories.")
	  ErrTranslationTemplateCatAlreadyHasKey = errors.New("Template already contains a key with this name.")
	  ErrTranslationTemplateCatDoesNotExist  = errors.New("Category does not exist.")
	  ErrTranslationTemplateCatAlreadyExist  = errors.New("Category already exists.")*/
)

// Language code constants.
const (
	LANGUAGE_ENGLISH           = "en"
	LANGUAGE_DUTCH             = "nl"
	LANGUAGE_GERMAN            = "de"
	LANGUAGE_FRENCH            = "fr"
	LANGUAGE_SPANISH           = "es"
	LANGUAGE_PORTUGUESE        = "pt_PT"
	LANGUAGE_PORTUGUESE_BRAZIL = "pt_BR"
)
