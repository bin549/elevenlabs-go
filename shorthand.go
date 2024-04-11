// Code generated by cmd/codegen/main.go; DO NOT EDIT.
// Run 'go generate' after adding new methods with a 'Client' pointer receiver.

package elevenlabs

import "io"

// TextToSpeech calls the TextToSpeech method on the default client.
func TextToSpeech(voiceID string, ttsReq TextToSpeechRequest, queries ...QueryFunc) ([]byte, error) {
	return getDefaultClient().TextToSpeech(voiceID, ttsReq, queries...)
}

// TextToSpeechStream calls the TextToSpeechStream method on the default client.
func TextToSpeechStream(streamWriter io.Writer, voiceID string, ttsReq TextToSpeechRequest, queries ...QueryFunc) error {
	return getDefaultClient().TextToSpeechStream(streamWriter, voiceID, ttsReq, queries...)
}

func SpeechToSpeech(voiceID string, stsReq SpeechToSpeechRequest, queries ...QueryFunc) ([]byte, error) {
	return getDefaultClient().SpeechToSpeech(voiceID, stsReq, queries...)
}

// GetModels calls the GetModels method on the default client.
func GetModels() ([]Model, error) {
	return getDefaultClient().GetModels()
}

// GetVoices calls the GetVoices method on the default client.
func GetVoices() ([]Voice, error) {
	return getDefaultClient().GetVoices()
}

// GetDefaultVoiceSettings calls the GetDefaultVoiceSettings method on the default client.
func GetDefaultVoiceSettings() (VoiceSettings, error) {
	return getDefaultClient().GetDefaultVoiceSettings()
}

// GetVoiceSettings calls the GetVoiceSettings method on the default client.
func GetVoiceSettings(voiceId string) (VoiceSettings, error) {
	return getDefaultClient().GetVoiceSettings(voiceId)
}

// GetVoice calls the GetVoice method on the default client.
func GetVoice(voiceId string, queries ...QueryFunc) (Voice, error) {
	return getDefaultClient().GetVoice(voiceId, queries...)
}

// DeleteVoice calls the DeleteVoice method on the default client.
func DeleteVoice(voiceId string) error {
	return getDefaultClient().DeleteVoice(voiceId)
}

// EditVoiceSettings calls the EditVoiceSettings method on the default client.
func EditVoiceSettings(voiceId string, settings VoiceSettings) error {
	return getDefaultClient().EditVoiceSettings(voiceId, settings)
}

// AddVoice calls the AddVoice method on the default client.
func AddVoice(voiceReq AddEditVoiceRequest) (string, error) {
	return getDefaultClient().AddVoice(voiceReq)
}

// EditVoice calls the EditVoice method on the default client.
func EditVoice(voiceId string, voiceReq AddEditVoiceRequest) error {
	return getDefaultClient().EditVoice(voiceId, voiceReq)
}

// DeleteSample calls the DeleteSample method on the default client.
func DeleteSample(voiceId, sampleId string) error {
	return getDefaultClient().DeleteSample(voiceId, sampleId)
}

// GetSampleAudio calls the GetSampleAudio method on the default client.
func GetSampleAudio(voiceId, sampleId string) ([]byte, error) {
	return getDefaultClient().GetSampleAudio(voiceId, sampleId)
}

// GetHistory calls the GetHistory method on the default client.
func GetHistory(queries ...QueryFunc) (GetHistoryResponse, NextHistoryPageFunc, error) {
	return getDefaultClient().GetHistory(queries...)
}

// GetHistoryItem calls the GetHistoryItem method on the default client.
func GetHistoryItem(itemId string) (HistoryItem, error) {
	return getDefaultClient().GetHistoryItem(itemId)
}

// DeleteHistoryItem calls the DeleteHistoryItem method on the default client.
func DeleteHistoryItem(itemId string) error {
	return getDefaultClient().DeleteHistoryItem(itemId)
}

// GetHistoryItemAudio calls the GetHistoryItemAudio method on the default client.
func GetHistoryItemAudio(itemId string) ([]byte, error) {
	return getDefaultClient().GetHistoryItemAudio(itemId)
}

// DownloadHistoryAudio calls the DownloadHistoryAudio method on the default client.
func DownloadHistoryAudio(dlReq DownloadHistoryRequest) ([]byte, error) {
	return getDefaultClient().DownloadHistoryAudio(dlReq)
}

// GetSubscription calls the GetSubscription method on the default client.
func GetSubscription() (Subscription, error) {
	return getDefaultClient().GetSubscription()
}

// GetUser calls the GetUser method on the default client.
func GetUser() (User, error) {
	return getDefaultClient().GetUser()
}
