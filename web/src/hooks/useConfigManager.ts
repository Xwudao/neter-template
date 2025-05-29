import { Config, initConfig, RawConfig } from '@/store/configState';
import useConfigState from '@/store/configState';
import { parseRawConfig } from '@/utils/configUtils';
import { useCallback, useEffect, useState } from 'react';

export const useConfigManager = () => {
  const {
    config: storeConfig,
    updateConfig: updateStoreConfig,
    resetConfig: resetStoreConfig,
    hasCache,
  } = useConfigState();

  const [configState, setConfigState] = useState<Config>(storeConfig);

  useEffect(() => {
    setConfigState(storeConfig);
  }, [storeConfig]);

  const updateConfig = useCallback(
    (config: Partial<Config>) => {
      const newConfig = { ...configState, ...config };
      setConfigState(newConfig);
      updateStoreConfig(config);
    },
    [configState, updateStoreConfig],
  );

  const updateRawConfig = useCallback(
    (rawConfig: Partial<RawConfig>) => {
      try {
        const parsedConfig = parseRawConfig(rawConfig);
        if (Object.keys(parsedConfig).length > 0) {
          updateStoreConfig(parsedConfig);
        }
      } catch (error) {
        console.error('Error updating raw config:', error);
      }
    },
    [updateStoreConfig],
  );

  const resetConfig = useCallback(() => {
    resetStoreConfig();
    const initialConfig = initConfig();
    setConfigState(initialConfig);
  }, [resetStoreConfig]);

  return {
    configState,
    updateConfig,
    updateRawConfig,
    resetConfig,
    hasCache,
  };
};
