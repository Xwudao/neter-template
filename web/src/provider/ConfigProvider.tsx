import { getApiGetCmnSiteConfig } from '@/api/siteConfigApi';
import PageLoading from '@/components/loading/PageLoading';
import { useConfigManager } from '@/hooks/useConfigManager';
import { Config, RawConfig } from '@/store/configState.ts';
import { useQuery } from '@tanstack/react-query';
import { createContext, ReactNode, useEffect } from 'react';

export interface ConfigContextType {
  config: Config;
  updateConfig: (config: Partial<Config>) => void;
  resetConfig: () => void;
  updateRawConfig: (rawConfig: Partial<RawConfig>) => void;
}

const ConfigContext = createContext<ConfigContextType>(null!);

const ConfigProvider = ({ children }: { children: ReactNode }) => {
  const { configState, updateConfig, updateRawConfig, resetConfig, hasCache } = useConfigManager();

  const {
    data: cmnConfig,
    error,
    isLoading,
  } = useQuery({
    queryKey: ['update-cmn-config'],
    queryFn: getApiGetCmnSiteConfig,
  });

  // 监听服务器配置数据变化，自动更新本地状态
  useEffect(() => {
    if (cmnConfig?.data) {
      updateRawConfig(cmnConfig.data);
    }
  }, [cmnConfig, updateRawConfig]);

  return (
    <ConfigContext.Provider
      value={{
        config: configState,
        updateConfig,
        resetConfig,
        updateRawConfig,
      }}>
      {error && <div className="text-red-500 text-center">Error loading configuration: {error.message}</div>}
      {isLoading && !hasCache() && <PageLoading text="Loading..." />}
      {(cmnConfig?.data || hasCache()) && children}
    </ConfigContext.Provider>
  );
};

export { ConfigContext };
export default ConfigProvider;
