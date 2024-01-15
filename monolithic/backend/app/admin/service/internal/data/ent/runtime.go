// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dict"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/dictdetail"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/menu"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/organization"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/position"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/role"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/schema"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	dictMixin := schema.Dict{}.Mixin()
	dictMixinFields0 := dictMixin[0].Fields()
	_ = dictMixinFields0
	dictFields := schema.Dict{}.Fields()
	_ = dictFields
	// dictDescID is the schema descriptor for id field.
	dictDescID := dictMixinFields0[0].Descriptor()
	// dict.IDValidator is a validator for the "id" field. It is called by the builders before save.
	dict.IDValidator = dictDescID.Validators[0].(func(uint32) error)
	dictdetailMixin := schema.DictDetail{}.Mixin()
	dictdetailMixinFields0 := dictdetailMixin[0].Fields()
	_ = dictdetailMixinFields0
	dictdetailFields := schema.DictDetail{}.Fields()
	_ = dictdetailFields
	// dictdetailDescDictID is the schema descriptor for dict_id field.
	dictdetailDescDictID := dictdetailFields[0].Descriptor()
	// dictdetail.DefaultDictID holds the default value on creation for the dict_id field.
	dictdetail.DefaultDictID = dictdetailDescDictID.Default.(uint32)
	// dictdetailDescOrderNo is the schema descriptor for order_no field.
	dictdetailDescOrderNo := dictdetailFields[1].Descriptor()
	// dictdetail.DefaultOrderNo holds the default value on creation for the order_no field.
	dictdetail.DefaultOrderNo = dictdetailDescOrderNo.Default.(int32)
	// dictdetailDescID is the schema descriptor for id field.
	dictdetailDescID := dictdetailMixinFields0[0].Descriptor()
	// dictdetail.IDValidator is a validator for the "id" field. It is called by the builders before save.
	dictdetail.IDValidator = dictdetailDescID.Validators[0].(func(uint32) error)
	menuMixin := schema.Menu{}.Mixin()
	menuMixinFields0 := menuMixin[0].Fields()
	_ = menuMixinFields0
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescOrderNo is the schema descriptor for order_no field.
	menuDescOrderNo := menuFields[2].Descriptor()
	// menu.DefaultOrderNo holds the default value on creation for the order_no field.
	menu.DefaultOrderNo = menuDescOrderNo.Default.(int32)
	// menuDescName is the schema descriptor for name field.
	menuDescName := menuFields[3].Descriptor()
	// menu.DefaultName holds the default value on creation for the name field.
	menu.DefaultName = menuDescName.Default.(string)
	// menu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menu.NameValidator = func() func(string) error {
		validators := menuDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// menuDescTitle is the schema descriptor for title field.
	menuDescTitle := menuFields[4].Descriptor()
	// menu.DefaultTitle holds the default value on creation for the title field.
	menu.DefaultTitle = menuDescTitle.Default.(string)
	// menu.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	menu.TitleValidator = menuDescTitle.Validators[0].(func(string) error)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[6].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[7].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescIcon is the schema descriptor for icon field.
	menuDescIcon := menuFields[8].Descriptor()
	// menu.DefaultIcon holds the default value on creation for the icon field.
	menu.DefaultIcon = menuDescIcon.Default.(string)
	// menu.IconValidator is a validator for the "icon" field. It is called by the builders before save.
	menu.IconValidator = menuDescIcon.Validators[0].(func(string) error)
	// menuDescIsExt is the schema descriptor for is_ext field.
	menuDescIsExt := menuFields[9].Descriptor()
	// menu.DefaultIsExt holds the default value on creation for the is_ext field.
	menu.DefaultIsExt = menuDescIsExt.Default.(bool)
	// menuDescExtURL is the schema descriptor for ext_url field.
	menuDescExtURL := menuFields[10].Descriptor()
	// menu.ExtURLValidator is a validator for the "ext_url" field. It is called by the builders before save.
	menu.ExtURLValidator = menuDescExtURL.Validators[0].(func(string) error)
	// menuDescKeepAlive is the schema descriptor for keep_alive field.
	menuDescKeepAlive := menuFields[14].Descriptor()
	// menu.DefaultKeepAlive holds the default value on creation for the keep_alive field.
	menu.DefaultKeepAlive = menuDescKeepAlive.Default.(bool)
	// menuDescShow is the schema descriptor for show field.
	menuDescShow := menuFields[15].Descriptor()
	// menu.DefaultShow holds the default value on creation for the show field.
	menu.DefaultShow = menuDescShow.Default.(bool)
	// menuDescHideTab is the schema descriptor for hide_tab field.
	menuDescHideTab := menuFields[16].Descriptor()
	// menu.DefaultHideTab holds the default value on creation for the hide_tab field.
	menu.DefaultHideTab = menuDescHideTab.Default.(bool)
	// menuDescHideMenu is the schema descriptor for hide_menu field.
	menuDescHideMenu := menuFields[17].Descriptor()
	// menu.DefaultHideMenu holds the default value on creation for the hide_menu field.
	menu.DefaultHideMenu = menuDescHideMenu.Default.(bool)
	// menuDescHideBreadcrumb is the schema descriptor for hide_breadcrumb field.
	menuDescHideBreadcrumb := menuFields[18].Descriptor()
	// menu.DefaultHideBreadcrumb holds the default value on creation for the hide_breadcrumb field.
	menu.DefaultHideBreadcrumb = menuDescHideBreadcrumb.Default.(bool)
	// menuDescID is the schema descriptor for id field.
	menuDescID := menuFields[0].Descriptor()
	// menu.IDValidator is a validator for the "id" field. It is called by the builders before save.
	menu.IDValidator = menuDescID.Validators[0].(func(int32) error)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationMixinFields2 := organizationMixin[2].Fields()
	_ = organizationMixinFields2
	organizationMixinFields4 := organizationMixin[4].Fields()
	_ = organizationMixinFields4
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescRemark is the schema descriptor for remark field.
	organizationDescRemark := organizationMixinFields4[0].Descriptor()
	// organization.DefaultRemark holds the default value on creation for the remark field.
	organization.DefaultRemark = organizationDescRemark.Default.(string)
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.DefaultName holds the default value on creation for the name field.
	organization.DefaultName = organizationDescName.Default.(string)
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
	// organizationDescOrderNo is the schema descriptor for order_no field.
	organizationDescOrderNo := organizationFields[2].Descriptor()
	// organization.DefaultOrderNo holds the default value on creation for the order_no field.
	organization.DefaultOrderNo = organizationDescOrderNo.Default.(int32)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields0[0].Descriptor()
	// organization.IDValidator is a validator for the "id" field. It is called by the builders before save.
	organization.IDValidator = organizationDescID.Validators[0].(func(uint32) error)
	positionMixin := schema.Position{}.Mixin()
	positionMixinFields0 := positionMixin[0].Fields()
	_ = positionMixinFields0
	positionMixinFields2 := positionMixin[2].Fields()
	_ = positionMixinFields2
	positionMixinFields4 := positionMixin[4].Fields()
	_ = positionMixinFields4
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescRemark is the schema descriptor for remark field.
	positionDescRemark := positionMixinFields4[0].Descriptor()
	// position.DefaultRemark holds the default value on creation for the remark field.
	position.DefaultRemark = positionDescRemark.Default.(string)
	// positionDescName is the schema descriptor for name field.
	positionDescName := positionFields[0].Descriptor()
	// position.DefaultName holds the default value on creation for the name field.
	position.DefaultName = positionDescName.Default.(string)
	// position.NameValidator is a validator for the "name" field. It is called by the builders before save.
	position.NameValidator = positionDescName.Validators[0].(func(string) error)
	// positionDescCode is the schema descriptor for code field.
	positionDescCode := positionFields[1].Descriptor()
	// position.DefaultCode holds the default value on creation for the code field.
	position.DefaultCode = positionDescCode.Default.(string)
	// position.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	position.CodeValidator = positionDescCode.Validators[0].(func(string) error)
	// positionDescParentID is the schema descriptor for parent_id field.
	positionDescParentID := positionFields[2].Descriptor()
	// position.DefaultParentID holds the default value on creation for the parent_id field.
	position.DefaultParentID = positionDescParentID.Default.(uint32)
	// positionDescOrderNo is the schema descriptor for order_no field.
	positionDescOrderNo := positionFields[3].Descriptor()
	// position.DefaultOrderNo holds the default value on creation for the order_no field.
	position.DefaultOrderNo = positionDescOrderNo.Default.(int32)
	// positionDescID is the schema descriptor for id field.
	positionDescID := positionMixinFields0[0].Descriptor()
	// position.IDValidator is a validator for the "id" field. It is called by the builders before save.
	position.IDValidator = positionDescID.Validators[0].(func(uint32) error)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleMixinFields2 := roleMixin[2].Fields()
	_ = roleMixinFields2
	roleMixinFields4 := roleMixin[4].Fields()
	_ = roleMixinFields4
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescRemark is the schema descriptor for remark field.
	roleDescRemark := roleMixinFields4[0].Descriptor()
	// role.DefaultRemark holds the default value on creation for the remark field.
	role.DefaultRemark = roleDescRemark.Default.(string)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCode is the schema descriptor for code field.
	roleDescCode := roleFields[1].Descriptor()
	// role.DefaultCode holds the default value on creation for the code field.
	role.DefaultCode = roleDescCode.Default.(string)
	// role.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	role.CodeValidator = roleDescCode.Validators[0].(func(string) error)
	// roleDescOrderNo is the schema descriptor for order_no field.
	roleDescOrderNo := roleFields[3].Descriptor()
	// role.DefaultOrderNo holds the default value on creation for the order_no field.
	role.DefaultOrderNo = roleDescOrderNo.Default.(int32)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.IDValidator is a validator for the "id" field. It is called by the builders before save.
	role.IDValidator = roleDescID.Validators[0].(func(uint32) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields3 := userMixin[3].Fields()
	_ = userMixinFields3
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[6].Descriptor()
	// user.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	user.NickNameValidator = userDescNickName.Validators[0].(func(string) error)
	// userDescRealName is the schema descriptor for real_name field.
	userDescRealName := userFields[7].Descriptor()
	// user.RealNameValidator is a validator for the "real_name" field. It is called by the builders before save.
	user.RealNameValidator = userDescRealName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[8].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[9].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[10].Descriptor()
	// user.AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	user.AvatarValidator = userDescAvatar.Validators[0].(func(string) error)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[12].Descriptor()
	// user.DefaultAddress holds the default value on creation for the address field.
	user.DefaultAddress = userDescAddress.Default.(string)
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescDescription is the schema descriptor for description field.
	userDescDescription := userFields[13].Descriptor()
	// user.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	user.DescriptionValidator = userDescDescription.Validators[0].(func(string) error)
	// userDescLastLoginIP is the schema descriptor for last_login_ip field.
	userDescLastLoginIP := userFields[16].Descriptor()
	// user.DefaultLastLoginIP holds the default value on creation for the last_login_ip field.
	user.DefaultLastLoginIP = userDescLastLoginIP.Default.(string)
	// user.LastLoginIPValidator is a validator for the "last_login_ip" field. It is called by the builders before save.
	user.LastLoginIPValidator = userDescLastLoginIP.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint32) error)
}
