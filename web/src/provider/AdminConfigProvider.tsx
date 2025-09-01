import { getApiGetSiteConfig, postApiUpdateSiteConfig } from '@/api/siteConfigApi';
import PageLoading from '@/components/loading/PageLoading';
import { useConfigManager } from '@/hooks/useConfigManager';
import { Config, RawConfig } from '@/store/configState.ts';
import { Toast } from '@douyinfe/semi-ui';
import { useMutation, useQuery } from '@tanstack/react-query';
import { createContext, ReactNode, useEffect } from 'react';

export interface AdminConfigContextType {
  config: Config;
  updating: boolean;

  handleUpdateConfig: (name: keyof Config, value: string | Record<string, any>, refresh?: boolean) => void;
  updateConfig: (config: Partial<Config>) => void;
  resetConfig: () => void;
  updateRawConfig: (rawConfig: Partial<RawConfig>) => void;
}

const AdminConfigContext = createContext<AdminConfigContextType>(null!);

const AdminConfigProvider = ({ children }: { children: ReactNode }) => {
  const { configState, updateConfig, updateRawConfig, resetConfig, hasCache } = useConfigManager();

  const {
    data: adminConfig,
    refetch: refetchAdminConfig,
    error,
    isLoading,
  } = useQuery({
    queryKey: ['admin-site-config'],
    queryFn: getApiGetSiteConfig,
  });

  // 监听服务器配置数据变化，自动更新本地状态
  useEffect(() => {
    if (adminConfig?.data) {
      updateRawConfig(adminConfig.data);
    }
  }, [adminConfig, updateRawConfig]);

  const { mutate, isPending } = useMutation({
    mutationKey: ['update-admin-config'],
    mutationFn: postApiUpdateSiteConfig,
  });

  const handleUpdateConfig = (name: keyof Config, value: string | Record<string, any>, isRefresh?: boolean) => {
    try {
      // 准备要发送到服务器的配置数据
      let configValue: string;
      if (typeof value === 'string') {
        configValue = value;
      } else {
        configValue = JSON.stringify(value);
      }

      // 发送到服务器
      mutate(
        { name, config: configValue },
        {
          onSuccess: () => {
            // 服务器更新成功后，更新本地状态
            updateConfig({ [name]: value });
            Toast.success('更新成功');
            if (isRefresh) {
              refetchAdminConfig();
            }
          },
          onError: (error) => {
            Toast.error(error.message || '更新失败');
            console.error(`Failed to update config field ${name}:`, error);
          },
        },
      );
    } catch (error) {
      console.error(`Error preparing config update for ${name}:`, error);
    }
  };

  return (
    <AdminConfigContext.Provider
      value={{
        config: configState,
        handleUpdateConfig,
        updateConfig,
        resetConfig,
        updating: isPending,
        updateRawConfig,
      }}>
      {error && <div className="text-red-500 text-center">Error loading configuration: {error.message}</div>}
      {isLoading && !hasCache() && <PageLoading text="Loading..." />}
      {(adminConfig?.data || hasCache()) && children}
    </AdminConfigContext.Provider>
  );
};

export { AdminConfigContext };
export default AdminConfigProvider;
