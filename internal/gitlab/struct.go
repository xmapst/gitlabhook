//go:build !windows

package gitlab

type UserInfo struct {
	ID                             int           `json:"id,omitempty"`
	Name                           string        `json:"name,omitempty"`
	Username                       string        `json:"username,omitempty"`
	State                          string        `json:"state,omitempty"`
	AvatarURL                      string        `json:"avatar_url,omitempty"`
	WebURL                         string        `json:"web_url,omitempty"`
	CreatedAt                      string        `json:"created_at,omitempty"`
	Bio                            string        `json:"bio,omitempty"`
	BioHTML                        string        `json:"bio_html,omitempty"`
	Location                       interface{}   `json:"location,omitempty"`
	PublicEmail                    string        `json:"public_email,omitempty"`
	Skype                          string        `json:"skype,omitempty"`
	Linkedin                       string        `json:"linkedin,omitempty"`
	Twitter                        string        `json:"twitter,omitempty"`
	WebsiteURL                     string        `json:"website_url,omitempty"`
	Organization                   interface{}   `json:"organization,omitempty"`
	JobTitle                       string        `json:"job_title,omitempty"`
	WorkInformation                interface{}   `json:"work_information,omitempty"`
	LastSignInAt                   string        `json:"last_sign_in_at,omitempty"`
	ConfirmedAt                    string        `json:"confirmed_at,omitempty"`
	LastActivityOn                 string        `json:"last_activity_on,omitempty"`
	Email                          string        `json:"email,omitempty"`
	ThemeID                        int           `json:"theme_id,omitempty"`
	ColorSchemeID                  int           `json:"color_scheme_id,omitempty"`
	ProjectsLimit                  int           `json:"projects_limit,omitempty"`
	CurrentSignInAt                string        `json:"current_sign_in_at,omitempty"`
	Identities                     []interface{} `json:"identities,omitempty"`
	CanCreateGroup                 bool          `json:"can_create_group,omitempty"`
	CanCreateProject               bool          `json:"can_create_project,omitempty"`
	TwoFactorEnabled               bool          `json:"two_factor_enabled,omitempty"`
	External                       bool          `json:"external,omitempty"`
	PrivateProfile                 bool          `json:"private_profile,omitempty"`
	SharedRunnersMinutesLimit      interface{}   `json:"shared_runners_minutes_limit,omitempty"`
	ExtraSharedRunnersMinutesLimit interface{}   `json:"extra_shared_runners_minutes_limit,omitempty"`
	IsAdmin                        bool          `json:"is_admin,omitempty"`
	Note                           interface{}   `json:"note,omitempty"`
	UsingLicenseSeat               bool          `json:"using_license_seat,omitempty"`
}
