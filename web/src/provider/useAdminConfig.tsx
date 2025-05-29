import { useContext } from 'react';
import { AdminConfigContext, AdminConfigContextType } from './AdminConfigProvider';

const useAdminConfig = (): AdminConfigContextType => {
  const context = useContext(AdminConfigContext);
  if (!context) {
    throw new Error('useAdminConfig must be used within an AdminConfigProvider');
  }
  return context;
};

export default useAdminConfig;
