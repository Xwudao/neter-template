import { useContext } from 'react';
import { ConfigContext } from '@/provider/ConfigProvider';

const useConfig = () => {
  return useContext(ConfigContext);
};

export default useConfig;
