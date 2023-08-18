/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AlternateKeyInitParameters struct {
}

type AlternateKeyObservation struct {
}

type AlternateKeyParameters struct {

	// The RSA parameter exponent.
	// +kubebuilder:validation:Optional
	RsaTokenKeyExponentSecretRef *v1.SecretKeySelector `json:"rsaTokenKeyExponentSecretRef,omitempty" tf:"-"`

	// The RSA parameter modulus.
	// +kubebuilder:validation:Optional
	RsaTokenKeyModulusSecretRef *v1.SecretKeySelector `json:"rsaTokenKeyModulusSecretRef,omitempty" tf:"-"`

	// The key value of the key. Specifies a symmetric key for token validation.
	// +kubebuilder:validation:Optional
	SymmetricTokenKeySecretRef *v1.SecretKeySelector `json:"symmetricTokenKeySecretRef,omitempty" tf:"-"`

	// The raw data field of a certificate in PKCS 12 format (X509Certificate2 in .NET). Specifies a certificate for token validation.
	// +kubebuilder:validation:Optional
	X509TokenKeyRawSecretRef *v1.SecretKeySelector `json:"x509TokenKeyRawSecretRef,omitempty" tf:"-"`
}

type ContentKeyPolicyInitParameters struct {

	// A description for the Policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// One or more policy_option blocks as defined below.
	PolicyOption []PolicyOptionInitParameters `json:"policyOption,omitempty" tf:"policy_option,omitempty"`
}

type ContentKeyPolicyObservation struct {

	// A description for the Policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Content Key Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// One or more policy_option blocks as defined below.
	PolicyOption []PolicyOptionObservation `json:"policyOption,omitempty" tf:"policy_option,omitempty"`

	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type ContentKeyPolicyParameters struct {

	// A description for the Policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Media Services account name. Changing this forces a new Content Key Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount in media to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount in media to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// One or more policy_option blocks as defined below.
	// +kubebuilder:validation:Optional
	PolicyOption []PolicyOptionParameters `json:"policyOption,omitempty" tf:"policy_option,omitempty"`

	// The name of the Resource Group where the Content Key Policy should exist. Changing this forces a new Content Key Policy to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type ExplicitAnalogTelevisionOutputRestrictionInitParameters struct {

	// Indicates whether this restriction is enforced on a best effort basis. Possible values are true or false. Defaults to false.
	BestEffortEnforced *bool `json:"bestEffortEnforced,omitempty" tf:"best_effort_enforced,omitempty"`

	// The restriction control bits. Possible value is integer between 0 and 3 inclusive.
	ControlBits *float64 `json:"controlBits,omitempty" tf:"control_bits,omitempty"`
}

type ExplicitAnalogTelevisionOutputRestrictionObservation struct {

	// Indicates whether this restriction is enforced on a best effort basis. Possible values are true or false. Defaults to false.
	BestEffortEnforced *bool `json:"bestEffortEnforced,omitempty" tf:"best_effort_enforced,omitempty"`

	// The restriction control bits. Possible value is integer between 0 and 3 inclusive.
	ControlBits *float64 `json:"controlBits,omitempty" tf:"control_bits,omitempty"`
}

type ExplicitAnalogTelevisionOutputRestrictionParameters struct {

	// Indicates whether this restriction is enforced on a best effort basis. Possible values are true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	BestEffortEnforced *bool `json:"bestEffortEnforced,omitempty" tf:"best_effort_enforced,omitempty"`

	// The restriction control bits. Possible value is integer between 0 and 3 inclusive.
	// +kubebuilder:validation:Optional
	ControlBits *float64 `json:"controlBits" tf:"control_bits,omitempty"`
}

type FairplayConfigurationInitParameters struct {

	// A offline_rental_configuration block as defined below.
	OfflineRentalConfiguration []OfflineRentalConfigurationInitParameters `json:"offlineRentalConfiguration,omitempty" tf:"offline_rental_configuration,omitempty"`

	// The rental and lease key type. Supported values are DualExpiry, PersistentLimited, PersistentUnlimited or Undefined.
	RentalAndLeaseKeyType *string `json:"rentalAndLeaseKeyType,omitempty" tf:"rental_and_lease_key_type,omitempty"`

	// The rental duration. Must be greater than 0.
	RentalDurationSeconds *float64 `json:"rentalDurationSeconds,omitempty" tf:"rental_duration_seconds,omitempty"`
}

type FairplayConfigurationObservation struct {

	// A offline_rental_configuration block as defined below.
	OfflineRentalConfiguration []OfflineRentalConfigurationObservation `json:"offlineRentalConfiguration,omitempty" tf:"offline_rental_configuration,omitempty"`

	// The rental and lease key type. Supported values are DualExpiry, PersistentLimited, PersistentUnlimited or Undefined.
	RentalAndLeaseKeyType *string `json:"rentalAndLeaseKeyType,omitempty" tf:"rental_and_lease_key_type,omitempty"`

	// The rental duration. Must be greater than 0.
	RentalDurationSeconds *float64 `json:"rentalDurationSeconds,omitempty" tf:"rental_duration_seconds,omitempty"`
}

type FairplayConfigurationParameters struct {

	// The key that must be used as FairPlay Application Secret key.
	// +kubebuilder:validation:Optional
	AskSecretRef *v1.SecretKeySelector `json:"askSecretRef,omitempty" tf:"-"`

	// A offline_rental_configuration block as defined below.
	// +kubebuilder:validation:Optional
	OfflineRentalConfiguration []OfflineRentalConfigurationParameters `json:"offlineRentalConfiguration,omitempty" tf:"offline_rental_configuration,omitempty"`

	// The password encrypting FairPlay certificate in PKCS 12 (pfx) format.
	// +kubebuilder:validation:Optional
	PfxPasswordSecretRef *v1.SecretKeySelector `json:"pfxPasswordSecretRef,omitempty" tf:"-"`

	// The Base64 representation of FairPlay certificate in PKCS 12 (pfx) format (including private key).
	// +kubebuilder:validation:Optional
	PfxSecretRef *v1.SecretKeySelector `json:"pfxSecretRef,omitempty" tf:"-"`

	// The rental and lease key type. Supported values are DualExpiry, PersistentLimited, PersistentUnlimited or Undefined.
	// +kubebuilder:validation:Optional
	RentalAndLeaseKeyType *string `json:"rentalAndLeaseKeyType,omitempty" tf:"rental_and_lease_key_type,omitempty"`

	// The rental duration. Must be greater than 0.
	// +kubebuilder:validation:Optional
	RentalDurationSeconds *float64 `json:"rentalDurationSeconds,omitempty" tf:"rental_duration_seconds,omitempty"`
}

type OfflineRentalConfigurationInitParameters struct {

	// Playback duration.
	PlaybackDurationSeconds *float64 `json:"playbackDurationSeconds,omitempty" tf:"playback_duration_seconds,omitempty"`

	// Storage duration.
	StorageDurationSeconds *float64 `json:"storageDurationSeconds,omitempty" tf:"storage_duration_seconds,omitempty"`
}

type OfflineRentalConfigurationObservation struct {

	// Playback duration.
	PlaybackDurationSeconds *float64 `json:"playbackDurationSeconds,omitempty" tf:"playback_duration_seconds,omitempty"`

	// Storage duration.
	StorageDurationSeconds *float64 `json:"storageDurationSeconds,omitempty" tf:"storage_duration_seconds,omitempty"`
}

type OfflineRentalConfigurationParameters struct {

	// Playback duration.
	// +kubebuilder:validation:Optional
	PlaybackDurationSeconds *float64 `json:"playbackDurationSeconds,omitempty" tf:"playback_duration_seconds,omitempty"`

	// Storage duration.
	// +kubebuilder:validation:Optional
	StorageDurationSeconds *float64 `json:"storageDurationSeconds,omitempty" tf:"storage_duration_seconds,omitempty"`
}

type PlayRightInitParameters struct {

	// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
	AgcAndColorStripeRestriction *float64 `json:"agcAndColorStripeRestriction,omitempty" tf:"agc_and_color_stripe_restriction,omitempty"`

	// Configures Unknown output handling settings of the license. Supported values are Allowed, AllowedWithVideoConstriction or NotAllowed.
	AllowPassingVideoContentToUnknownOutput *string `json:"allowPassingVideoContentToUnknownOutput,omitempty" tf:"allow_passing_video_content_to_unknown_output,omitempty"`

	// Specifies the output protection level for compressed digital audio. Supported values are 100, 150 or 200.
	AnalogVideoOpl *float64 `json:"analogVideoOpl,omitempty" tf:"analog_video_opl,omitempty"`

	// Specifies the output protection level for compressed digital audio.Supported values are 100, 150, 200, 250 or 300.
	CompressedDigitalAudioOpl *float64 `json:"compressedDigitalAudioOpl,omitempty" tf:"compressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for compressed digital video. Supported values are 400 or 500.
	CompressedDigitalVideoOpl *float64 `json:"compressedDigitalVideoOpl,omitempty" tf:"compressed_digital_video_opl,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	DigitalVideoOnlyContentRestriction *bool `json:"digitalVideoOnlyContentRestriction,omitempty" tf:"digital_video_only_content_restriction,omitempty"`

	// An explicit_analog_television_output_restriction block as defined above.
	ExplicitAnalogTelevisionOutputRestriction []ExplicitAnalogTelevisionOutputRestrictionInitParameters `json:"explicitAnalogTelevisionOutputRestriction,omitempty" tf:"explicit_analog_television_output_restriction,omitempty"`

	// The amount of time that the license is valid after the license is first used to play content.
	FirstPlayExpiration *string `json:"firstPlayExpiration,omitempty" tf:"first_play_expiration,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	ImageConstraintForAnalogComponentVideoRestriction *bool `json:"imageConstraintForAnalogComponentVideoRestriction,omitempty" tf:"image_constraint_for_analog_component_video_restriction,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	ImageConstraintForAnalogComputerMonitorRestriction *bool `json:"imageConstraintForAnalogComputerMonitorRestriction,omitempty" tf:"image_constraint_for_analog_computer_monitor_restriction,omitempty"`

	// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
	ScmsRestriction *float64 `json:"scmsRestriction,omitempty" tf:"scms_restriction,omitempty"`

	// Specifies the output protection level for uncompressed digital audio. Supported values are 100, 150, 200, 250 or 300.
	UncompressedDigitalAudioOpl *float64 `json:"uncompressedDigitalAudioOpl,omitempty" tf:"uncompressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for uncompressed digital video. Supported values are 100, 250, 270 or 300.
	UncompressedDigitalVideoOpl *float64 `json:"uncompressedDigitalVideoOpl,omitempty" tf:"uncompressed_digital_video_opl,omitempty"`
}

type PlayRightObservation struct {

	// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
	AgcAndColorStripeRestriction *float64 `json:"agcAndColorStripeRestriction,omitempty" tf:"agc_and_color_stripe_restriction,omitempty"`

	// Configures Unknown output handling settings of the license. Supported values are Allowed, AllowedWithVideoConstriction or NotAllowed.
	AllowPassingVideoContentToUnknownOutput *string `json:"allowPassingVideoContentToUnknownOutput,omitempty" tf:"allow_passing_video_content_to_unknown_output,omitempty"`

	// Specifies the output protection level for compressed digital audio. Supported values are 100, 150 or 200.
	AnalogVideoOpl *float64 `json:"analogVideoOpl,omitempty" tf:"analog_video_opl,omitempty"`

	// Specifies the output protection level for compressed digital audio.Supported values are 100, 150, 200, 250 or 300.
	CompressedDigitalAudioOpl *float64 `json:"compressedDigitalAudioOpl,omitempty" tf:"compressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for compressed digital video. Supported values are 400 or 500.
	CompressedDigitalVideoOpl *float64 `json:"compressedDigitalVideoOpl,omitempty" tf:"compressed_digital_video_opl,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	DigitalVideoOnlyContentRestriction *bool `json:"digitalVideoOnlyContentRestriction,omitempty" tf:"digital_video_only_content_restriction,omitempty"`

	// An explicit_analog_television_output_restriction block as defined above.
	ExplicitAnalogTelevisionOutputRestriction []ExplicitAnalogTelevisionOutputRestrictionObservation `json:"explicitAnalogTelevisionOutputRestriction,omitempty" tf:"explicit_analog_television_output_restriction,omitempty"`

	// The amount of time that the license is valid after the license is first used to play content.
	FirstPlayExpiration *string `json:"firstPlayExpiration,omitempty" tf:"first_play_expiration,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	ImageConstraintForAnalogComponentVideoRestriction *bool `json:"imageConstraintForAnalogComponentVideoRestriction,omitempty" tf:"image_constraint_for_analog_component_video_restriction,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	ImageConstraintForAnalogComputerMonitorRestriction *bool `json:"imageConstraintForAnalogComputerMonitorRestriction,omitempty" tf:"image_constraint_for_analog_computer_monitor_restriction,omitempty"`

	// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
	ScmsRestriction *float64 `json:"scmsRestriction,omitempty" tf:"scms_restriction,omitempty"`

	// Specifies the output protection level for uncompressed digital audio. Supported values are 100, 150, 200, 250 or 300.
	UncompressedDigitalAudioOpl *float64 `json:"uncompressedDigitalAudioOpl,omitempty" tf:"uncompressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for uncompressed digital video. Supported values are 100, 250, 270 or 300.
	UncompressedDigitalVideoOpl *float64 `json:"uncompressedDigitalVideoOpl,omitempty" tf:"uncompressed_digital_video_opl,omitempty"`
}

type PlayRightParameters struct {

	// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
	// +kubebuilder:validation:Optional
	AgcAndColorStripeRestriction *float64 `json:"agcAndColorStripeRestriction,omitempty" tf:"agc_and_color_stripe_restriction,omitempty"`

	// Configures Unknown output handling settings of the license. Supported values are Allowed, AllowedWithVideoConstriction or NotAllowed.
	// +kubebuilder:validation:Optional
	AllowPassingVideoContentToUnknownOutput *string `json:"allowPassingVideoContentToUnknownOutput,omitempty" tf:"allow_passing_video_content_to_unknown_output,omitempty"`

	// Specifies the output protection level for compressed digital audio. Supported values are 100, 150 or 200.
	// +kubebuilder:validation:Optional
	AnalogVideoOpl *float64 `json:"analogVideoOpl,omitempty" tf:"analog_video_opl,omitempty"`

	// Specifies the output protection level for compressed digital audio.Supported values are 100, 150, 200, 250 or 300.
	// +kubebuilder:validation:Optional
	CompressedDigitalAudioOpl *float64 `json:"compressedDigitalAudioOpl,omitempty" tf:"compressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for compressed digital video. Supported values are 400 or 500.
	// +kubebuilder:validation:Optional
	CompressedDigitalVideoOpl *float64 `json:"compressedDigitalVideoOpl,omitempty" tf:"compressed_digital_video_opl,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	// +kubebuilder:validation:Optional
	DigitalVideoOnlyContentRestriction *bool `json:"digitalVideoOnlyContentRestriction,omitempty" tf:"digital_video_only_content_restriction,omitempty"`

	// An explicit_analog_television_output_restriction block as defined above.
	// +kubebuilder:validation:Optional
	ExplicitAnalogTelevisionOutputRestriction []ExplicitAnalogTelevisionOutputRestrictionParameters `json:"explicitAnalogTelevisionOutputRestriction,omitempty" tf:"explicit_analog_television_output_restriction,omitempty"`

	// The amount of time that the license is valid after the license is first used to play content.
	// +kubebuilder:validation:Optional
	FirstPlayExpiration *string `json:"firstPlayExpiration,omitempty" tf:"first_play_expiration,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	// +kubebuilder:validation:Optional
	ImageConstraintForAnalogComponentVideoRestriction *bool `json:"imageConstraintForAnalogComponentVideoRestriction,omitempty" tf:"image_constraint_for_analog_component_video_restriction,omitempty"`

	// Enables the Image Constraint For Analog Component Video Restriction in the license.
	// +kubebuilder:validation:Optional
	ImageConstraintForAnalogComputerMonitorRestriction *bool `json:"imageConstraintForAnalogComputerMonitorRestriction,omitempty" tf:"image_constraint_for_analog_computer_monitor_restriction,omitempty"`

	// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
	// +kubebuilder:validation:Optional
	ScmsRestriction *float64 `json:"scmsRestriction,omitempty" tf:"scms_restriction,omitempty"`

	// Specifies the output protection level for uncompressed digital audio. Supported values are 100, 150, 200, 250 or 300.
	// +kubebuilder:validation:Optional
	UncompressedDigitalAudioOpl *float64 `json:"uncompressedDigitalAudioOpl,omitempty" tf:"uncompressed_digital_audio_opl,omitempty"`

	// Specifies the output protection level for uncompressed digital video. Supported values are 100, 250, 270 or 300.
	// +kubebuilder:validation:Optional
	UncompressedDigitalVideoOpl *float64 `json:"uncompressedDigitalVideoOpl,omitempty" tf:"uncompressed_digital_video_opl,omitempty"`
}

type PlayreadyConfigurationLicenseInitParameters struct {

	// A flag indicating whether test devices can use the license.
	AllowTestDevices *bool `json:"allowTestDevices,omitempty" tf:"allow_test_devices,omitempty"`

	// The begin date of license.
	BeginDate *string `json:"beginDate,omitempty" tf:"begin_date,omitempty"`

	// Specifies that the content key ID is in the PlayReady header.
	ContentKeyLocationFromHeaderEnabled *bool `json:"contentKeyLocationFromHeaderEnabled,omitempty" tf:"content_key_location_from_header_enabled,omitempty"`

	// The content key ID. Specifies that the content key ID is specified in the PlayReady configuration.
	ContentKeyLocationFromKeyID *string `json:"contentKeyLocationFromKeyId,omitempty" tf:"content_key_location_from_key_id,omitempty"`

	// The PlayReady content type. Supported values are UltraVioletDownload, UltraVioletStreaming or Unspecified.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The expiration date of license.
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The license type. Supported values are NonPersistent or Persistent.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A play_right block as defined above.
	PlayRight []PlayRightInitParameters `json:"playRight,omitempty" tf:"play_right,omitempty"`

	// The relative begin date of license.
	RelativeBeginDate *string `json:"relativeBeginDate,omitempty" tf:"relative_begin_date,omitempty"`

	// The relative expiration date of license.
	RelativeExpirationDate *string `json:"relativeExpirationDate,omitempty" tf:"relative_expiration_date,omitempty"`

	// The security level of the PlayReady license. Possible values are SL150, SL2000 and SL3000. Please see this document for more information about security level. See this document for more information about SL3000 support.
	SecurityLevel *string `json:"securityLevel,omitempty" tf:"security_level,omitempty"`
}

type PlayreadyConfigurationLicenseObservation struct {

	// A flag indicating whether test devices can use the license.
	AllowTestDevices *bool `json:"allowTestDevices,omitempty" tf:"allow_test_devices,omitempty"`

	// The begin date of license.
	BeginDate *string `json:"beginDate,omitempty" tf:"begin_date,omitempty"`

	// Specifies that the content key ID is in the PlayReady header.
	ContentKeyLocationFromHeaderEnabled *bool `json:"contentKeyLocationFromHeaderEnabled,omitempty" tf:"content_key_location_from_header_enabled,omitempty"`

	// The content key ID. Specifies that the content key ID is specified in the PlayReady configuration.
	ContentKeyLocationFromKeyID *string `json:"contentKeyLocationFromKeyId,omitempty" tf:"content_key_location_from_key_id,omitempty"`

	// The PlayReady content type. Supported values are UltraVioletDownload, UltraVioletStreaming or Unspecified.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The expiration date of license.
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The license type. Supported values are NonPersistent or Persistent.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A play_right block as defined above.
	PlayRight []PlayRightObservation `json:"playRight,omitempty" tf:"play_right,omitempty"`

	// The relative begin date of license.
	RelativeBeginDate *string `json:"relativeBeginDate,omitempty" tf:"relative_begin_date,omitempty"`

	// The relative expiration date of license.
	RelativeExpirationDate *string `json:"relativeExpirationDate,omitempty" tf:"relative_expiration_date,omitempty"`

	// The security level of the PlayReady license. Possible values are SL150, SL2000 and SL3000. Please see this document for more information about security level. See this document for more information about SL3000 support.
	SecurityLevel *string `json:"securityLevel,omitempty" tf:"security_level,omitempty"`
}

type PlayreadyConfigurationLicenseParameters struct {

	// A flag indicating whether test devices can use the license.
	// +kubebuilder:validation:Optional
	AllowTestDevices *bool `json:"allowTestDevices,omitempty" tf:"allow_test_devices,omitempty"`

	// The begin date of license.
	// +kubebuilder:validation:Optional
	BeginDate *string `json:"beginDate,omitempty" tf:"begin_date,omitempty"`

	// Specifies that the content key ID is in the PlayReady header.
	// +kubebuilder:validation:Optional
	ContentKeyLocationFromHeaderEnabled *bool `json:"contentKeyLocationFromHeaderEnabled,omitempty" tf:"content_key_location_from_header_enabled,omitempty"`

	// The content key ID. Specifies that the content key ID is specified in the PlayReady configuration.
	// +kubebuilder:validation:Optional
	ContentKeyLocationFromKeyID *string `json:"contentKeyLocationFromKeyId,omitempty" tf:"content_key_location_from_key_id,omitempty"`

	// The PlayReady content type. Supported values are UltraVioletDownload, UltraVioletStreaming or Unspecified.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The expiration date of license.
	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// The grace period of license.
	// +kubebuilder:validation:Optional
	GracePeriodSecretRef *v1.SecretKeySelector `json:"gracePeriodSecretRef,omitempty" tf:"-"`

	// The license type. Supported values are NonPersistent or Persistent.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// A play_right block as defined above.
	// +kubebuilder:validation:Optional
	PlayRight []PlayRightParameters `json:"playRight,omitempty" tf:"play_right,omitempty"`

	// The relative begin date of license.
	// +kubebuilder:validation:Optional
	RelativeBeginDate *string `json:"relativeBeginDate,omitempty" tf:"relative_begin_date,omitempty"`

	// The relative expiration date of license.
	// +kubebuilder:validation:Optional
	RelativeExpirationDate *string `json:"relativeExpirationDate,omitempty" tf:"relative_expiration_date,omitempty"`

	// The security level of the PlayReady license. Possible values are SL150, SL2000 and SL3000. Please see this document for more information about security level. See this document for more information about SL3000 support.
	// +kubebuilder:validation:Optional
	SecurityLevel *string `json:"securityLevel,omitempty" tf:"security_level,omitempty"`
}

type PolicyOptionInitParameters struct {

	// Enable a configuration for non-DRM keys.
	ClearKeyConfigurationEnabled *bool `json:"clearKeyConfigurationEnabled,omitempty" tf:"clear_key_configuration_enabled,omitempty"`

	// A fairplay_configuration block as defined above. Check license requirements here https://docs.microsoft.com/azure/media-services/latest/fairplay-license-overview.
	FairplayConfiguration []FairplayConfigurationInitParameters `json:"fairplayConfiguration,omitempty" tf:"fairplay_configuration,omitempty"`

	// The name which should be used for this Policy Option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enable an open restriction. License or key will be delivered on every request.
	OpenRestrictionEnabled *bool `json:"openRestrictionEnabled,omitempty" tf:"open_restriction_enabled,omitempty"`

	// One or more playready_configuration_license blocks as defined above.
	PlayreadyConfigurationLicense []PlayreadyConfigurationLicenseInitParameters `json:"playreadyConfigurationLicense,omitempty" tf:"playready_configuration_license,omitempty"`

	// The custom response data of the PlayReady configuration. This only applies when playready_configuration_license is specified.
	PlayreadyResponseCustomData *string `json:"playreadyResponseCustomData,omitempty" tf:"playready_response_custom_data,omitempty"`

	// A token_restriction block as defined below.
	TokenRestriction []TokenRestrictionInitParameters `json:"tokenRestriction,omitempty" tf:"token_restriction,omitempty"`

	// The Widevine template.
	WidevineConfigurationTemplate *string `json:"widevineConfigurationTemplate,omitempty" tf:"widevine_configuration_template,omitempty"`
}

type PolicyOptionObservation struct {

	// Enable a configuration for non-DRM keys.
	ClearKeyConfigurationEnabled *bool `json:"clearKeyConfigurationEnabled,omitempty" tf:"clear_key_configuration_enabled,omitempty"`

	// A fairplay_configuration block as defined above. Check license requirements here https://docs.microsoft.com/azure/media-services/latest/fairplay-license-overview.
	FairplayConfiguration []FairplayConfigurationObservation `json:"fairplayConfiguration,omitempty" tf:"fairplay_configuration,omitempty"`

	// The name which should be used for this Policy Option.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Enable an open restriction. License or key will be delivered on every request.
	OpenRestrictionEnabled *bool `json:"openRestrictionEnabled,omitempty" tf:"open_restriction_enabled,omitempty"`

	// One or more playready_configuration_license blocks as defined above.
	PlayreadyConfigurationLicense []PlayreadyConfigurationLicenseObservation `json:"playreadyConfigurationLicense,omitempty" tf:"playready_configuration_license,omitempty"`

	// The custom response data of the PlayReady configuration. This only applies when playready_configuration_license is specified.
	PlayreadyResponseCustomData *string `json:"playreadyResponseCustomData,omitempty" tf:"playready_response_custom_data,omitempty"`

	// A token_restriction block as defined below.
	TokenRestriction []TokenRestrictionObservation `json:"tokenRestriction,omitempty" tf:"token_restriction,omitempty"`

	// The Widevine template.
	WidevineConfigurationTemplate *string `json:"widevineConfigurationTemplate,omitempty" tf:"widevine_configuration_template,omitempty"`
}

type PolicyOptionParameters struct {

	// Enable a configuration for non-DRM keys.
	// +kubebuilder:validation:Optional
	ClearKeyConfigurationEnabled *bool `json:"clearKeyConfigurationEnabled,omitempty" tf:"clear_key_configuration_enabled,omitempty"`

	// A fairplay_configuration block as defined above. Check license requirements here https://docs.microsoft.com/azure/media-services/latest/fairplay-license-overview.
	// +kubebuilder:validation:Optional
	FairplayConfiguration []FairplayConfigurationParameters `json:"fairplayConfiguration,omitempty" tf:"fairplay_configuration,omitempty"`

	// The name which should be used for this Policy Option.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Enable an open restriction. License or key will be delivered on every request.
	// +kubebuilder:validation:Optional
	OpenRestrictionEnabled *bool `json:"openRestrictionEnabled,omitempty" tf:"open_restriction_enabled,omitempty"`

	// One or more playready_configuration_license blocks as defined above.
	// +kubebuilder:validation:Optional
	PlayreadyConfigurationLicense []PlayreadyConfigurationLicenseParameters `json:"playreadyConfigurationLicense,omitempty" tf:"playready_configuration_license,omitempty"`

	// The custom response data of the PlayReady configuration. This only applies when playready_configuration_license is specified.
	// +kubebuilder:validation:Optional
	PlayreadyResponseCustomData *string `json:"playreadyResponseCustomData,omitempty" tf:"playready_response_custom_data,omitempty"`

	// A token_restriction block as defined below.
	// +kubebuilder:validation:Optional
	TokenRestriction []TokenRestrictionParameters `json:"tokenRestriction,omitempty" tf:"token_restriction,omitempty"`

	// The Widevine template.
	// +kubebuilder:validation:Optional
	WidevineConfigurationTemplate *string `json:"widevineConfigurationTemplate,omitempty" tf:"widevine_configuration_template,omitempty"`
}

type RequiredClaimInitParameters struct {

	// Token claim type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Token claim value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RequiredClaimObservation struct {

	// Token claim type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Token claim value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RequiredClaimParameters struct {

	// Token claim type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Token claim value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TokenRestrictionInitParameters struct {

	// One or more alternate_key block as defined above.
	AlternateKey []AlternateKeyInitParameters `json:"alternateKey,omitempty" tf:"alternate_key,omitempty"`

	// The audience for the token.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The token issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The OpenID connect discovery document.
	OpenIDConnectDiscoveryDocument *string `json:"openIdConnectDiscoveryDocument,omitempty" tf:"open_id_connect_discovery_document,omitempty"`

	// One or more required_claim blocks as defined above.
	RequiredClaim []RequiredClaimInitParameters `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// The type of token. Supported values are Jwt or Swt.
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type TokenRestrictionObservation struct {

	// One or more alternate_key block as defined above.
	AlternateKey []AlternateKeyParameters `json:"alternateKey,omitempty" tf:"alternate_key,omitempty"`

	// The audience for the token.
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The token issuer.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The OpenID connect discovery document.
	OpenIDConnectDiscoveryDocument *string `json:"openIdConnectDiscoveryDocument,omitempty" tf:"open_id_connect_discovery_document,omitempty"`

	// One or more required_claim blocks as defined above.
	RequiredClaim []RequiredClaimObservation `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// The type of token. Supported values are Jwt or Swt.
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

type TokenRestrictionParameters struct {

	// One or more alternate_key block as defined above.
	// +kubebuilder:validation:Optional
	AlternateKey []AlternateKeyParameters `json:"alternateKey,omitempty" tf:"alternate_key,omitempty"`

	// The audience for the token.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The token issuer.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The OpenID connect discovery document.
	// +kubebuilder:validation:Optional
	OpenIDConnectDiscoveryDocument *string `json:"openIdConnectDiscoveryDocument,omitempty" tf:"open_id_connect_discovery_document,omitempty"`

	// The RSA parameter exponent.
	// +kubebuilder:validation:Optional
	PrimaryRsaTokenKeyExponentSecretRef *v1.SecretKeySelector `json:"primaryRsaTokenKeyExponentSecretRef,omitempty" tf:"-"`

	// The RSA parameter modulus.
	// +kubebuilder:validation:Optional
	PrimaryRsaTokenKeyModulusSecretRef *v1.SecretKeySelector `json:"primaryRsaTokenKeyModulusSecretRef,omitempty" tf:"-"`

	// The key value of the key. Specifies a symmetric key for token validation.
	// +kubebuilder:validation:Optional
	PrimarySymmetricTokenKeySecretRef *v1.SecretKeySelector `json:"primarySymmetricTokenKeySecretRef,omitempty" tf:"-"`

	// The raw data field of a certificate in PKCS 12 format (X509Certificate2 in .NET). Specifies a certificate for token validation.
	// +kubebuilder:validation:Optional
	PrimaryX509TokenKeyRawSecretRef *v1.SecretKeySelector `json:"primaryX509TokenKeyRawSecretRef,omitempty" tf:"-"`

	// One or more required_claim blocks as defined above.
	// +kubebuilder:validation:Optional
	RequiredClaim []RequiredClaimParameters `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// The type of token. Supported values are Jwt or Swt.
	// +kubebuilder:validation:Optional
	TokenType *string `json:"tokenType,omitempty" tf:"token_type,omitempty"`
}

// ContentKeyPolicySpec defines the desired state of ContentKeyPolicy
type ContentKeyPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContentKeyPolicyParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ContentKeyPolicyInitParameters `json:"initProvider,omitempty"`
}

// ContentKeyPolicyStatus defines the observed state of ContentKeyPolicy.
type ContentKeyPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContentKeyPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContentKeyPolicy is the Schema for the ContentKeyPolicys API. Manages a Content Key Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContentKeyPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyOption) || has(self.initProvider.policyOption)",message="policyOption is a required parameter"
	Spec   ContentKeyPolicySpec   `json:"spec"`
	Status ContentKeyPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentKeyPolicyList contains a list of ContentKeyPolicys
type ContentKeyPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentKeyPolicy `json:"items"`
}

// Repository type metadata.
var (
	ContentKeyPolicy_Kind             = "ContentKeyPolicy"
	ContentKeyPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContentKeyPolicy_Kind}.String()
	ContentKeyPolicy_KindAPIVersion   = ContentKeyPolicy_Kind + "." + CRDGroupVersion.String()
	ContentKeyPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ContentKeyPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ContentKeyPolicy{}, &ContentKeyPolicyList{})
}
