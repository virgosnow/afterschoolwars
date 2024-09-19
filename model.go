package main

type Response struct {
	Error struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		ReloginFlag  int    `json:"relogin_flag"`
	} `json:"error"`
	Data struct {
		UserBasic struct {
			FamilyName        string `json:"family_name"`
			FirstName         string `json:"first_name"`
			Live              string `json:"live"`
			Level             string `json:"level"`
			Exp               string `json:"exp"`
			ExpProgress       int    `json:"exp_progress"`
			RainbowStoneNum   int    `json:"rainbow_stone_num"`
			GoldStoneNum      string `json:"gold_stone_num"`
			SilverStoneNum    string `json:"silver_stone_num"`
			TutorialProgress  string `json:"tutorial_progress"`
			MaxCost           string `json:"max_cost"`
			NewMissionNum     int    `json:"new_mission_num"`
			NewPresentNum     int    `json:"new_present_num"`
			NewNoticeNum      int    `json:"new_notice_num"`
			NewHaremNum       int    `json:"new_harem_num"`
			MstCardId         string `json:"mst_card_id"`
			MstCharacterId    string `json:"mst_character_id"`
			MstCardHash       string `json:"mst_card_hash"`
			MstCharacterHash  string `json:"mst_character_hash"`
			TeamNum           string `json:"team_num"`
			AccessoriesMax    string `json:"accessories_max"`
			XtensionAccessory string `json:"xtension_accessory"`
			WarehouseMax      string `json:"warehouse_max"`
			XtensionWarehouse string `json:"xtension_warehouse"`
			Ap                struct {
				ApMax         string `json:"ap_max"`
				ApNow         string `json:"ap_now"`
				ApRecoverAt   int    `json:"ap_recover_at"`
				ApRecoverTime int    `json:"ap_recover_time"`
			} `json:"ap"`
			Bp struct {
				BpMax         string `json:"bp_max"`
				BpNow         string `json:"bp_now"`
				BpRecoverAt   int    `json:"bp_recover_at"`
				BpRecoverTime int    `json:"bp_recover_time"`
			} `json:"bp"`
			Tp struct {
				TpMax         string `json:"tp_max"`
				TpNow         string `json:"tp_now"`
				TpRecoverAt   int    `json:"tp_recover_at"`
				TpRecoverTime int    `json:"tp_recover_time"`
			} `json:"tp"`
			Time string `json:"time"`
		} `json:"user_basic"`
		HaremScene []struct {
			TabType      int    `json:"tab_type"`
			MstSceneId   string `json:"mst_scene_id"`
			MstSceneHash string `json:"mst_scene_hash"`
			Name         string `json:"name"`
			IsNew        string `json:"is_new"`
		} `json:"harem_scene"`
	} `json:"data"`
	ServerTime string `json:"server_time"`
}

type EVResponse struct {
	Error struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		ReloginFlag  int    `json:"relogin_flag"`
	} `json:"error"`
	Data struct {
		User struct {
			UserId               int    `json:"user_id"`
			Fps                  int    `json:"fps"`
			Url                  string `json:"url"`
			ErrorUrl             string `json:"error_url"`
			IsDebug              bool   `json:"is_debug"`
			IsPc                 bool   `json:"is_pc"`
			VoiceMute            bool   `json:"voice_mute"`
			VoiceVolume          int    `json:"voice_volume"`
			SeMute               bool   `json:"se_mute"`
			SeVolume             int    `json:"se_volume"`
			BgmMute              bool   `json:"bgm_mute"`
			BgmVolume            int    `json:"bgm_volume"`
			ScrWordsTransparency int    `json:"scr_words_transparency"`
			ScrWordsSpeed        int    `json:"scr_words_speed"`
			ScrWordsShadow       int    `json:"scr_words_shadow"`
			ScrAutoWait          int    `json:"scr_auto_wait"`
			ScrAutoVoice         int    `json:"scr_auto_voice"`
			ScrAnimation         int    `json:"scr_animation"`
			LivelyGacha          int    `json:"lively_gacha"`
			LivelyCompose        int    `json:"lively_compose"`
			LivelyBattle         int    `json:"lively_battle"`
		} `json:"user"`
		Code struct {
			Images     []string      `json:"images"`
			Sounds     []string      `json:"sounds"`
			Articles   []string      `json:"articles"`
			Characters []interface{} `json:"characters"`
			Tags       []string      `json:"tags"`
			Parameters [][]string    `json:"parameters"`
			Commands   [][]string    `json:"commands"`
		} `json:"code"`
	} `json:"data"`
	ServerTime string `json:"server_time"`
}
