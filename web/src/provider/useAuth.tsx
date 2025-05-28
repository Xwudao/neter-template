import { AuthContext } from '@/provider/AuthProvider';

const useAuth = () => {
  return useContext(AuthContext);
};

export default useAuth;
