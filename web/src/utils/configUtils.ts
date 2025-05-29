import { Config, initConfig, RawConfig } from '@/store/configState';

// 通用 JSON 解析器
export const parseJsonField = <T,>(jsonStr: string | null | undefined, defaultValue: T): T => {
  if (!jsonStr || typeof jsonStr !== 'string') {
    return defaultValue;
  }
  try {
    const parsed = JSON.parse(jsonStr);
    return typeof parsed === 'object' && parsed !== null ? { ...defaultValue, ...parsed } : defaultValue;
  } catch (error) {
    console.warn('Failed to parse JSON field:', error);
    return defaultValue;
  }
};

// 将原始配置转换为解析后的配置
export const parseRawConfig = (rawConfig: Partial<RawConfig>): Partial<Config> => {
  if (!rawConfig || typeof rawConfig !== 'object') {
    return {};
  }

  const defaultConfig = initConfig();
  const parsedConfig: Partial<Config> = {};

  try {
    if (rawConfig.site_info !== undefined) {
      parsedConfig.site_info = parseJsonField(rawConfig.site_info, defaultConfig.site_info);
    }

    if (rawConfig.site_seo !== undefined) {
      parsedConfig.site_seo = parseJsonField(rawConfig.site_seo, defaultConfig.site_seo);
    }
  } catch (error) {
    console.error('Error parsing raw config:', error);
    return {};
  }

  return parsedConfig;
};
