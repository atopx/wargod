package conf

import (
	"fmt"
	"github.com/atopx/clever"
	"github.com/goccy/go-json"
	"os"
	"os/user"
	"path/filepath"
	"wargod/types"
)

type Config struct {
	AutoAccept    bool          `json:"auto_accept"`    // 自动应答
	AutoStatus    bool          `json:"auto_status"`    // 自动设置段位状态
	StatusContent StatusContent `json:"status_content"` // 自动设置段位内容
}

type StatusContent struct {
	StatusMessage        string             `json:"status_message"`
	RankedLeagueQueue    types.RankedQueue  `json:"ranked_league_queue"`
	RankedLeagueTier     types.Tier         `json:"ranked_league_tier"`
	RankedLeagueDivision types.TierDivision `json:"ranked_league_division"`
}

func (sc StatusContent) Data() []byte {
	data := fmt.Sprintf(`{"statusMessage":"%s","lol":{"rankedLeagueQueue":"%s","rankedLeagueTier":"%s", "rankedLeagueDivision": "%s"}}`,
		sc.StatusMessage,
		sc.RankedLeagueQueue,
		sc.RankedLeagueTier,
		sc.RankedLeagueDivision,
	)
	return clever.Bytes(data)
}

var Entry *Config

func configPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	return filepath.Join(usr.HomeDir, "AppData", "Roaming", "wargod", "config.json"), nil
}

func LoadConfig() error {
	// 配置文件路径
	path, err := configPath()
	if err != nil {
		return err
	}

	// 读取配置文件
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，初始化一个默认配置并写入文件
			Entry = &Config{}
			return SaveConfig()
		}
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置文件内容
	err = json.Unmarshal(data, &Entry)
	if err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}

func SaveConfig() error {
	// 配置文件路径
	path, err := configPath()
	if err != nil {
		return err
	}

	// 确保配置目录存在
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// 序列化配置为JSON
	data, err := json.MarshalIndent(Entry, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// 写入配置文件
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
