package tts

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type EdgeTTSService struct {
	expiredAt int64
	endpoint  *Endpoint
	clientID  string
}

type Endpoint struct {
	Region string `json:"r"`
	Token  string `json:"t"`
}

type TTSRequest struct {
	Text   string `json:"text"`
	Voice  string `json:"voice"`
	Rate   int    `json:"rate"`
	Pitch  int    `json:"pitch"`
	Format string `json:"format"`
}

type Voice struct {
	ShortName       string `json:"ShortName"`
	LocalName       string `json:"LocalName"`
	Locale          string `json:"Locale"`
	Gender          string `json:"Gender"`
	WordsPerMinute  int    `json:"WordsPerMinute"`
	SampleRateHertz string `json:"SampleRateHertz"`
}

func NewEdgeTTSService() *EdgeTTSService {
	return &EdgeTTSService{
		clientID: uuid.New().String(),
	}
}

func (s *EdgeTTSService) GenerateAudio(text, voice string, rate, pitch int, format string) ([]byte, error) {
	if err := s.refreshEndpoint(); err != nil {
		return nil, err
	}

	ssml := s.generateSSML(text, voice, rate, pitch)
	regionUrl := fmt.Sprintf("https://%s.tts.speech.microsoft.com/cognitiveservices/v1", s.endpoint.Region)

	req, err := http.NewRequest("POST", regionUrl, strings.NewReader(ssml))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", s.endpoint.Token)
	req.Header.Set("Content-Type", "application/ssml+xml")
	req.Header.Set("X-Microsoft-OutputFormat", format)
	req.Header.Set("User-Agent", "okhttp/4.5.0")
	req.Header.Set("Origin", "https://azure.microsoft.com")
	req.Header.Set("Referer", "https://azure.microsoft.com/")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TTS 请求失败，状态码 %d", resp.StatusCode)
	}

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return audioData, nil
}

func (s *EdgeTTSService) GetVoices(localeFilter string) ([]Voice, error) {
	voices, err := s.getVoiceList()
	if err != nil {
		return nil, err
	}

	if localeFilter != "" {
		var filtered []Voice
		for _, voice := range voices {
			if strings.Contains(strings.ToLower(voice.Locale), strings.ToLower(localeFilter)) {
				filtered = append(filtered, voice)
			}
		}
		voices = filtered
	}

	return voices, nil
}

func (s *EdgeTTSService) FormatVoicesAsMap(voices []Voice) map[string]string {
	voiceMap := make(map[string]string)
	for _, voice := range voices {
		voiceMap[voice.ShortName] = voice.LocalName
	}
	return voiceMap
}

func (s *EdgeTTSService) FormatVoicesAsText(voices []Voice) string {
	var result []string
	for _, voice := range voices {
		result = append(result, s.formatVoiceItem(voice))
	}
	return strings.Join(result, "\n")
}

func (s *EdgeTTSService) generateSSML(text, voice string, rate, pitch int) string {
	return fmt.Sprintf(`<speak xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="http://www.w3.org/2001/mstts" version="1.0" xml:lang="zh-CN">
    <voice name="%s">
        <mstts:express-as style="general" styledegree="1.0" role="default">
            <prosody rate="%d%%" pitch="%d%%" volume="50">%s</prosody>
        </mstts:express-as>
    </voice>
</speak>`, voice, rate, pitch, text)
}

func (s *EdgeTTSService) formatVoiceItem(voice Voice) string {
	gender := "1"
	if voice.Gender == "Female" {
		gender = "0"
	}

	sampleRate := voice.SampleRateHertz
	if sampleRate == "" {
		sampleRate = "24000"
	}

	return fmt.Sprintf(`
- !!org.nobody.multitts.tts.speaker.Speaker
  avatar: ''
  code: %s
  desc: ''
  extendUI: ''
  gender: %s
  name: %s
  note: 'wpm: %d'
  param: ''
  sampleRate: %s
  speed: 1.5
  type: 1
  volume: 1`, voice.ShortName, gender, voice.LocalName, voice.WordsPerMinute, sampleRate)
}

func (s *EdgeTTSService) getVoiceList() ([]Voice, error) {
	req, err := http.NewRequest("GET", "https://eastus.api.speech.microsoft.com/cognitiveservices/voices/list", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("X-Ms-Useragent", "SpeechStudio/2021.05.001")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://azure.microsoft.com")
	req.Header.Set("Referer", "https://azure.microsoft.com")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("获取语音列表失败，状态码 %d", resp.StatusCode)
	}

	var voices []Voice
	if err := json.NewDecoder(resp.Body).Decode(&voices); err != nil {
		return nil, err
	}

	return voices, nil
}

func (s *EdgeTTSService) refreshEndpoint() error {
	now := time.Now().Unix()
	if s.expiredAt == 0 || now > s.expiredAt-60 {
		endpoint, err := s.getEndpoint()
		if err != nil {
			return err
		}

		s.endpoint = endpoint

		// Parse JWT token to get expiry time
		token, _, err := new(jwt.Parser).ParseUnverified(endpoint.Token, jwt.MapClaims{})
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if exp, ok := claims["exp"].(float64); ok {
					s.expiredAt = int64(exp)
				}
			}
		}

		if s.expiredAt == 0 {
			s.expiredAt = now + 3600 // Default 1 hour
		}

		s.clientID = uuid.New().String()
		fmt.Printf("获取 Endpoint, 过期时间剩余: %.2f 分钟\n", float64(s.expiredAt-now)/60)
	} else {
		fmt.Printf("过期时间剩余: %.2f 分钟\n", float64(s.expiredAt-now)/60)
	}

	return nil
}

func (s *EdgeTTSService) getEndpoint() (*Endpoint, error) {
	endpointURL := "https://dev.microsofttranslator.com/apps/endpoint?api-version=1.0"

	signature, err := s.generateSignature(endpointURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpointURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "zh-Hans")
	req.Header.Set("X-ClientVersion", "4.0.530a 5fe1dc6c")
	req.Header.Set("X-UserId", "0f04d16a175c411e")
	req.Header.Set("X-HomeGeographicRegion", "zh-Hans-CN")
	req.Header.Set("X-ClientTraceId", s.clientID)
	req.Header.Set("X-MT-Signature", signature)
	req.Header.Set("User-Agent", "okhttp/4.5.0")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept-Encoding", "gzip")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("获取 Endpoint 失败，状态码 %d", resp.StatusCode)
	}

	var endpoint Endpoint
	if err := json.NewDecoder(resp.Body).Decode(&endpoint); err != nil {
		return nil, err
	}

	return &endpoint, nil
}

func (s *EdgeTTSService) generateSignature(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	encodedURL := url.QueryEscape(parsedURL.Host + parsedURL.Path + "?" + parsedURL.RawQuery)
	uuidStr := s.generateUUID()
	formattedDate := s.formatDate()
	bytesToSign := fmt.Sprintf("MSTranslatorAndroidApp%s%s%s", encodedURL, formattedDate, uuidStr)
	bytesToSign = strings.ToLower(bytesToSign)

	key, err := base64.StdEncoding.DecodeString("oik6PdDdMnOXemTbwvMn9de/h9lFnfBaCWbGMMZqqoSaQaqUOqjVGm5NqsmjcBI1x+sS9ugjB55HEJWRiFXYFw==")
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(bytesToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("MSTranslatorAndroidApp::%s::%s::%s", signature, formattedDate, uuidStr), nil
}

func (s *EdgeTTSService) formatDate() string {
	return strings.ToLower(time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05") + " GMT")
}

func (s *EdgeTTSService) generateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
