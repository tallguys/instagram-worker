package account

type EdgeOwnerToTimelineMedia struct {
	Count    int `json:"count"`
	PageInfo struct {
		HasNextPage bool   `json:"has_next_page"`
		EndCursor   string `json:"end_cursor"`
	} `json:"page_info"`
	Edges []struct {
		Node struct {
			Typename           string `json:"__typename"`
			ID                 string `json:"id"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						Text string `json:"text"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			Shortcode          string `json:"shortcode"`
			EdgeMediaToComment struct {
				Count int `json:"count"`
			} `json:"edge_media_to_comment"`
			CommentsDisabled bool `json:"comments_disabled"`
			TakenAtTimestamp int  `json:"taken_at_timestamp"`
			Dimensions       struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			DisplayURL  string `json:"display_url"`
			EdgeLikedBy struct {
				Count int `json:"count"`
			} `json:"edge_liked_by"`
			EdgeMediaPreviewLike struct {
				Count int `json:"count"`
			} `json:"edge_media_preview_like"`
			Location struct {
				ID            string `json:"id"`
				HasPublicPage bool   `json:"has_public_page"`
				Name          string `json:"name"`
				Slug          string `json:"slug"`
			} `json:"location"`
			GatingInfo             interface{} `json:"gating_info"`
			FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
			FactCheckInformation   interface{} `json:"fact_check_information"`
			MediaPreview           interface{} `json:"media_preview"`
			Owner                  struct {
				ID       string `json:"id"`
				Username string `json:"username"`
			} `json:"owner"`
			ThumbnailSrc       string `json:"thumbnail_src"`
			ThumbnailResources []struct {
				Src          string `json:"src"`
				ConfigWidth  int    `json:"config_width"`
				ConfigHeight int    `json:"config_height"`
			} `json:"thumbnail_resources"`
			IsVideo              bool   `json:"is_video"`
			AccessibilityCaption string `json:"accessibility_caption"`
		} `json:"node"`
	} `json:"edges"`
}

type UserProfile struct {
	Biography              string      `json:"biography"`
	BlockedByViewer        bool        `json:"blocked_by_viewer"`
	RestrictedByViewer     interface{} `json:"restricted_by_viewer"`
	CountryBlock           bool        `json:"country_block"`
	ExternalURL            string      `json:"external_url"`
	ExternalURLLinkshimmed string      `json:"external_url_linkshimmed"`
	EdgeFollowedBy         struct {
		Count int `json:"count"`
	} `json:"edge_followed_by"`
	FollowedByViewer bool `json:"followed_by_viewer"`
	EdgeFollow       struct {
		Count int `json:"count"`
	} `json:"edge_follow"`
	FollowsViewer        bool        `json:"follows_viewer"`
	FullName             string      `json:"full_name"`
	HasArEffects         bool        `json:"has_ar_effects"`
	HasChannel           bool        `json:"has_channel"`
	HasBlockedViewer     bool        `json:"has_blocked_viewer"`
	HighlightReelCount   int         `json:"highlight_reel_count"`
	HasRequestedViewer   bool        `json:"has_requested_viewer"`
	ID                   string      `json:"id"`
	IsBusinessAccount    bool        `json:"is_business_account"`
	IsJoinedRecently     bool        `json:"is_joined_recently"`
	BusinessCategoryName string      `json:"business_category_name"`
	CategoryID           string      `json:"category_id"`
	OverallCategoryName  interface{} `json:"overall_category_name"`
	IsPrivate            bool        `json:"is_private"`
	IsVerified           bool        `json:"is_verified"`
	EdgeMutualFollowedBy struct {
		Count int           `json:"count"`
		Edges []interface{} `json:"edges"`
	} `json:"edge_mutual_followed_by"`
	ProfilePicURL          string      `json:"profile_pic_url"`
	ProfilePicURLHd        string      `json:"profile_pic_url_hd"`
	RequestedByViewer      bool        `json:"requested_by_viewer"`
	Username               string      `json:"username"`
	ConnectedFbPage        interface{} `json:"connected_fb_page"`
	EdgeFelixVideoTimeline struct {
		Count    int `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
		Edges []struct {
			Node struct {
				Typename           string `json:"__typename"`
				ID                 string `json:"id"`
				EdgeMediaToCaption struct {
					Edges []struct {
						Node struct {
							Text string `json:"text"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"edge_media_to_caption"`
				Shortcode          string `json:"shortcode"`
				EdgeMediaToComment struct {
					Count int `json:"count"`
				} `json:"edge_media_to_comment"`
				CommentsDisabled bool `json:"comments_disabled"`
				TakenAtTimestamp int  `json:"taken_at_timestamp"`
				Dimensions       struct {
					Height int `json:"height"`
					Width  int `json:"width"`
				} `json:"dimensions"`
				DisplayURL  string `json:"display_url"`
				EdgeLikedBy struct {
					Count int `json:"count"`
				} `json:"edge_liked_by"`
				EdgeMediaPreviewLike struct {
					Count int `json:"count"`
				} `json:"edge_media_preview_like"`
				Location               interface{} `json:"location"`
				GatingInfo             interface{} `json:"gating_info"`
				FactCheckOverallRating interface{} `json:"fact_check_overall_rating"`
				FactCheckInformation   interface{} `json:"fact_check_information"`
				MediaPreview           string      `json:"media_preview"`
				Owner                  struct {
					ID       string `json:"id"`
					Username string `json:"username"`
				} `json:"owner"`
				ThumbnailSrc       string `json:"thumbnail_src"`
				ThumbnailResources []struct {
					Src          string `json:"src"`
					ConfigWidth  int    `json:"config_width"`
					ConfigHeight int    `json:"config_height"`
				} `json:"thumbnail_resources"`
				IsVideo              bool `json:"is_video"`
				FelixProfileGridCrop struct {
					CropLeft   float64 `json:"crop_left"`
					CropRight  float64 `json:"crop_right"`
					CropTop    float64 `json:"crop_top"`
					CropBottom float64 `json:"crop_bottom"`
				} `json:"felix_profile_grid_crop"`
				EncodingStatus interface{} `json:"encoding_status"`
				IsPublished    bool        `json:"is_published"`
				ProductType    string      `json:"product_type"`
				Title          string      `json:"title"`
				VideoDuration  float64     `json:"video_duration"`
				VideoViewCount int         `json:"video_view_count"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_felix_video_timeline"`
	EdgeOwnerToTimelineMedia EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
	EdgeSavedMedia           struct {
		Count    int `json:"count"`
		PageInfo struct {
			HasNextPage bool        `json:"has_next_page"`
			EndCursor   interface{} `json:"end_cursor"`
		} `json:"page_info"`
		Edges []interface{} `json:"edges"`
	} `json:"edge_saved_media"`
	EdgeMediaCollections struct {
		Count    int `json:"count"`
		PageInfo struct {
			HasNextPage bool        `json:"has_next_page"`
			EndCursor   interface{} `json:"end_cursor"`
		} `json:"page_info"`
		Edges []interface{} `json:"edges"`
	} `json:"edge_media_collections"`
}

// SharedData user's shared data
type SharedData struct {
	Config struct {
		CsrfToken string      `json:"csrf_token"`
		Viewer    interface{} `json:"viewer"`
		ViewerID  interface{} `json:"viewerId"`
	} `json:"config"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string `json:"logging_page_id"`
			ShowSuggestedProfiles bool   `json:"show_suggested_profiles"`
			ShowFollowDialog      bool   `json:"show_follow_dialog"`
			Graphql               struct {
				User UserProfile `json:"user"`
			} `json:"graphql"`
			ToastContentOnLoad interface{} `json:"toast_content_on_load"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Hostname              string  `json:"hostname"`
	IsWhitelistedCrawlBot bool    `json:"is_whitelisted_crawl_bot"`
	DeploymentStage       string  `json:"deployment_stage"`
	Platform              string  `json:"platform"`
	Nonce                 string  `json:"nonce"`
	MidPct                float64 `json:"mid_pct"`
	ZeroData              struct {
	} `json:"zero_data"`
	CacheSchemaVersion int `json:"cache_schema_version"`
	ServerChecks       struct {
	} `json:"server_checks"`
	Knobx struct {
		Num4  bool `json:"4"`
		Num17 bool `json:"17"`
		Num20 bool `json:"20"`
		Num22 bool `json:"22"`
		Num23 bool `json:"23"`
		Num24 bool `json:"24"`
		Num25 bool `json:"25"`
	} `json:"knobx"`
	ToCache struct {
		Gatekeepers struct {
			Num4   bool `json:"4"`
			Num5   bool `json:"5"`
			Num6   bool `json:"6"`
			Num7   bool `json:"7"`
			Num8   bool `json:"8"`
			Num9   bool `json:"9"`
			Num10  bool `json:"10"`
			Num11  bool `json:"11"`
			Num12  bool `json:"12"`
			Num13  bool `json:"13"`
			Num14  bool `json:"14"`
			Num15  bool `json:"15"`
			Num16  bool `json:"16"`
			Num18  bool `json:"18"`
			Num19  bool `json:"19"`
			Num23  bool `json:"23"`
			Num24  bool `json:"24"`
			Num26  bool `json:"26"`
			Num27  bool `json:"27"`
			Num28  bool `json:"28"`
			Num29  bool `json:"29"`
			Num31  bool `json:"31"`
			Num32  bool `json:"32"`
			Num34  bool `json:"34"`
			Num35  bool `json:"35"`
			Num38  bool `json:"38"`
			Num40  bool `json:"40"`
			Num41  bool `json:"41"`
			Num43  bool `json:"43"`
			Num59  bool `json:"59"`
			Num61  bool `json:"61"`
			Num62  bool `json:"62"`
			Num63  bool `json:"63"`
			Num64  bool `json:"64"`
			Num65  bool `json:"65"`
			Num67  bool `json:"67"`
			Num68  bool `json:"68"`
			Num69  bool `json:"69"`
			Num71  bool `json:"71"`
			Num72  bool `json:"72"`
			Num73  bool `json:"73"`
			Num74  bool `json:"74"`
			Num75  bool `json:"75"`
			Num76  bool `json:"76"`
			Num77  bool `json:"77"`
			Num78  bool `json:"78"`
			Num79  bool `json:"79"`
			Num81  bool `json:"81"`
			Num82  bool `json:"82"`
			Num83  bool `json:"83"`
			Num84  bool `json:"84"`
			Num86  bool `json:"86"`
			Num88  bool `json:"88"`
			Num91  bool `json:"91"`
			Num93  bool `json:"93"`
			Num95  bool `json:"95"`
			Num96  bool `json:"96"`
			Num97  bool `json:"97"`
			Num98  bool `json:"98"`
			Num99  bool `json:"99"`
			Num100 bool `json:"100"`
		} `json:"gatekeepers"`
		Qe struct {
			Num0 struct {
				P struct {
					Num4 bool `json:"4"`
					Num7 bool `json:"7"`
					Num8 bool `json:"8"`
					Num9 bool `json:"9"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"0"`
			Num2 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"2"`
			Num4 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"4"`
			Num5 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"5"`
			Num6 struct {
				P struct {
					Num1  bool `json:"1"`
					Num5  bool `json:"5"`
					Num6  bool `json:"6"`
					Num7  bool `json:"7"`
					Num9  bool `json:"9"`
					Num10 bool `json:"10"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"6"`
			Num10 struct {
				P struct {
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"10"`
			Num12 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"12"`
			Num13 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"13"`
			Num16 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"16"`
			Num17 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"17"`
			Num19 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"19"`
			Num21 struct {
				P struct {
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"21"`
			Num22 struct {
				P struct {
					Num1  bool    `json:"1"`
					Num2  float64 `json:"2"`
					Num3  float64 `json:"3"`
					Num4  float64 `json:"4"`
					Num10 float64 `json:"10"`
					Num11 int     `json:"11"`
					Num12 int     `json:"12"`
					Num13 bool    `json:"13"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"22"`
			Num23 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"23"`
			Num25 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"25"`
			Num26 struct {
				P struct {
					Num0 string `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"26"`
			Num28 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"28"`
			Num29 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"29"`
			Num30 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"30"`
			Num31 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"31"`
			Num33 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"33"`
			Num34 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"34"`
			Num35 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"35"`
			Num36 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"36"`
			Num37 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"37"`
			Num39 struct {
				P struct {
					Num0  bool `json:"0"`
					Num6  bool `json:"6"`
					Num7  bool `json:"7"`
					Num8  bool `json:"8"`
					Num10 bool `json:"10"`
					Num11 bool `json:"11"`
					Num14 bool `json:"14"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"39"`
			Num41 struct {
				P struct {
					Num3 bool `json:"3"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"41"`
			Num42 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"42"`
			Num43 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"43"`
			Num44 struct {
				P struct {
					Num1 string  `json:"1"`
					Num2 float64 `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"44"`
			Num45 struct {
				P struct {
					Num12 bool   `json:"12"`
					Num13 bool   `json:"13"`
					Num17 int    `json:"17"`
					Num18 bool   `json:"18"`
					Num19 int    `json:"19"`
					Num22 bool   `json:"22"`
					Num23 string `json:"23"`
					Num24 bool   `json:"24"`
					Num25 string `json:"25"`
					Num26 string `json:"26"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"45"`
			Num46 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"46"`
			Num47 struct {
				P struct {
					Num0  bool `json:"0"`
					Num1  bool `json:"1"`
					Num2  bool `json:"2"`
					Num3  bool `json:"3"`
					Num4  bool `json:"4"`
					Num6  bool `json:"6"`
					Num8  bool `json:"8"`
					Num9  bool `json:"9"`
					Num10 bool `json:"10"`
					Num11 bool `json:"11"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"47"`
			Num49 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"49"`
			Num50 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"50"`
			Num53 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"53"`
			Num54 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"54"`
			Num55 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"55"`
			Num56 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"56"`
			Num58 struct {
				P struct {
					Num0 float64 `json:"0"`
					Num1 bool    `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"58"`
			Num59 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"59"`
			Num62 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"62"`
			Num64 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"64"`
			Num65 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"65"`
			Num66 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"66"`
			Num67 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
					Num5 bool `json:"5"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"67"`
			Num68 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"68"`
			Num69 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"69"`
			Num70 struct {
				P struct {
					Num1 string `json:"1"`
					Num2 string `json:"2"`
					Num3 string `json:"3"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"70"`
			Num71 struct {
				P struct {
					Num1 string `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"71"`
			Num72 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"72"`
			Num73 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"73"`
			Num74 struct {
				P struct {
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
					Num5 bool `json:"5"`
					Num6 bool `json:"6"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"74"`
			Num75 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"75"`
			Num76 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"76"`
			Num77 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"77"`
			Num78 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"78"`
			Num80 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"80"`
			Num81 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"81"`
			Num82 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"82"`
			Num84 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"84"`
			Num85 struct {
				P struct {
					Num0 bool   `json:"0"`
					Num1 string `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"85"`
			Num86 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"86"`
			Num87 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"87"`
			Num89 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"89"`
			Num90 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"90"`
			Num91 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"91"`
			Num92 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"92"`
			Num93 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"93"`
			Num95 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"95"`
			Num96 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"96"`
			AppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"app_upsell"`
			IglAppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"igl_app_upsell"`
			Notif struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"notif"`
			Onetaplogin struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"onetaplogin"`
			MultiregIter struct {
				G string `json:"g"`
				P struct {
					HasPrioritizedPhone string `json:"has_prioritized_phone"`
				} `json:"p"`
			} `json:"multireg_iter"`
			FelixClearFbCookie struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_clear_fb_cookie"`
			FelixCreationDurationLimits struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_duration_limits"`
			FelixCreationFbCrossposting struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting"`
			FelixCreationFbCrosspostingV2 struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting_v2"`
			FelixCreationValidation struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"felix_creation_validation"`
			MwebTopicalExplore struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"mweb_topical_explore"`
			PostOptions struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"post_options"`
			Iglscioi struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglscioi"`
			StickerTray struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"sticker_tray"`
			WebSentry struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"web_sentry"`
		} `json:"qe"`
		ProbablyHasApp bool `json:"probably_has_app"`
		Cb             bool `json:"cb"`
	} `json:"to_cache"`
	DeviceID   string `json:"device_id"`
	Encryption struct {
		KeyID     string `json:"key_id"`
		PublicKey string `json:"public_key"`
	} `json:"encryption"`
	RolloutHash   string `json:"rollout_hash"`
	BundleVariant string `json:"bundle_variant"`
	IsCanary      bool   `json:"is_canary"`
}
