package ipa

import (
	"encoding/json"
)

// HbacRecord encapsulates hbac data returned from ipa user commands
type HbacRecord struct {
	Dn            string   `json:"dn"`
	Description   []string `json:"description"`
	Enable        []string `json:"ipaenabledflag"`
	Hosts         []string `json:"memberhost_host"`
	HostGroups    []string `json:"memberhost_hostgroup"`
	Services      []string `json:"memberservice_hbacsvc"`
	ServicesGroup []string `json:"memberservice_hbacsvcgroup"`
	Users         []string `json:"memberhost_user"`
}

// Fetch hbac details by call the FreeIPA user-show method
func (c *Client) HbacShow(uid string) (*HbacRecord, error) {

	options := map[string]interface{}{
		"no_members": false,
		"all":        true}

	res, err := c.rpc("hbacrule_show", []string{uid}, options)

	if err != nil {
		return nil, err
	}

	var hbacRec HbacRecord
	err = json.Unmarshal(res.Result.Data, &hbacRec)
	if err != nil {
		return nil, err
	}

	return &hbacRec, nil
}

// // Update ssh public keys for user uid. Returns the fingerprints on success.
// func (c *Client) UpdateSSHPubKeys(uid string, keys []string) ([]string, error) {
// 	options := map[string]interface{}{
// 		"no_members":   false,
// 		"ipasshpubkey": keys,
// 		"all":          false}

// 	res, err := c.rpc("user_mod", []string{uid}, options)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var hbacRec HbacRecord
// 	err = json.Unmarshal(res.Result.Data, &hbacRec)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return hbacRec.SSHPubKeyFps, nil
// }

// // Update mobile number. Currently will store only a single number. Any
// // existing numbers will be overwritten.
// func (c *Client) UpdateMobileNumber(uid string, number string) error {
// 	options := map[string]interface{}{
// 		"no_members": false,
// 		"mobile":     []string{number},
// 		"all":        false}

// 	_, err := c.rpc("user_mod", []string{uid}, options)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Reset user password and return new random password
// func (c *Client) ResetPassword(uid string) (string, error) {

// 	options := map[string]interface{}{
// 		"no_members": false,
// 		"random":     true,
// 		"all":        true}

// 	res, err := c.rpc("user_mod", []string{uid}, options)

// 	if err != nil {
// 		return "", err
// 	}

// 	var hbacRec HbacRecord
// 	err = json.Unmarshal(res.Result.Data, &hbacRec)
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(hbacRec.Randompassword) == 0 {
// 		return "", errors.New("ipa: failed to reset user password. empty random password returned")
// 	}

// 	return hbacRec.Randompassword, nil
// }

// // Change user password. This will run the passwd ipa command. Optionally
// // provide an OTP if required
// func (c *Client) ChangePassword(uid, old_passwd, new_passwd, otpcode string) error {

// 	options := map[string]interface{}{
// 		"current_password": old_passwd,
// 		"password":         new_passwd,
// 	}

// 	if len(otpcode) > 0 {
// 		options["otp"] = otpcode
// 	}

// 	_, err := c.rpc("passwd", []string{uid}, options)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Set user password. In FreeIPA when a password is first set or when a
// // password is later reset it is marked as immediately expired and requires the
// // owner to perform a password change. See here
// // https://www.freeipa.org/page/New_Passwords_Expired for more details. This
// // function exists to circumvent the "new passwords expired" feature of FreeIPA
// // and allow an administrator to set a new password for a user without it being
// // expired. This is acheived, for example, by first calling ResetPassword()
// // then immediately calling this function. *WARNING* See
// // https://www.freeipa.org/page/Self-Service_Password_Reset for security issues
// // and possible weaknesses of this approach.
// func (c *Client) SetPassword(uid, old_passwd, new_passwd, otpcode string) error {
// 	ipaUrl := fmt.Sprintf("https://%s/ipa/session/change_password", c.host)

// 	form := url.Values{
// 		"user":         {uid},
// 		"otp":          {otpcode},
// 		"old_password": {old_passwd},
// 		"new_password": {new_passwd}}
// 	req, err := http.NewRequest("POST", ipaUrl, strings.NewReader(form.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Referer", fmt.Sprintf("https://%s/ipa", c.host))

// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{RootCAs: ipaCertPool}}
// 	client := &http.Client{Transport: tr}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != 200 {
// 		return fmt.Errorf("ipa: change password failed with HTTP status code: %d", res.StatusCode)
// 	}

// 	status := res.Header.Get("x-ipa-pwchange-result")
// 	if status == "policy-error" {
// 		return &ErrPasswordPolicy{}
// 	} else if status == "invalid-password" {
// 		return &ErrInvalidPassword{}
// 	} else if strings.ToLower(status) != "ok" {
// 		return errors.New("ipa: change password failed. Unknown status")
// 	}

// 	return nil
// }

// // Update user authentication types.
// func (c *Client) SetAuthTypes(uid string, types []string) error {
// 	options := map[string]interface{}{
// 		"no_members":      false,
// 		"ipauserauthtype": types,
// 		"all":             false}

// 	if len(types) == 0 {
// 		options["ipauserauthtype"] = ""
// 	}

// 	_, err := c.rpc("user_mod", []string{uid}, options)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Disable User Account
// func (c *Client) UserDisable(uid string) error {
// 	_, err := c.rpc("user_disable", []string{uid}, nil)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Enable User Account
// func (c *Client) UserEnable(uid string) error {
// 	_, err := c.rpc("user_enable", []string{uid}, nil)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // Add new user. If random is true a random password will be created for the
// // user. Note this requires "User Administrators" Privilege in FreeIPA.
// func (c *Client) UserAdd(uid, email, first, last, homedir, shell string, random bool) (*HbacRecord, error) {
// 	var options = map[string]interface{}{
// 		"mail":      email,
// 		"givenname": first,
// 		"sn":        last}

// 	if len(homedir) > 0 {
// 		options["homedirectory"] = homedir
// 	}

// 	if len(shell) > 0 {
// 		options["loginshell"] = shell
// 	}

// 	if random {
// 		options["random"] = true
// 	}

// 	res, err := c.rpc("user_add", []string{uid}, options)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var hbacRec HbacRecord
// 	err = json.Unmarshal(res.Result.Data, &hbacRec)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &hbacRec, nil
// }
