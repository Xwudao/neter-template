import { getApiUserInfo, User } from '@/api/userApi.ts';
import { KEY_TOKEN, UserRole } from '@/core/constants.ts';
import useUserState from '@/store/userState.ts';
import { useQuery } from '@tanstack/react-query';
import Cookies from 'js-cookie';
import { createContext, ReactNode } from 'react';

export interface AuthContextType {
  toLogin: (user: User, ok?: () => void) => void;
  user: User;
  logged: boolean;
  isAdmin: boolean;
  logout: (onOk?: () => void) => void;
}
const AuthContext = createContext<AuthContextType>(null!);

const AuthProvider = ({ children }: { children: ReactNode }) => {
  const { userInfo, resetInfo, updateUser } = useUserState();
  const [userState, setUserState] = useState<User>(userInfo);
  const storageToken = localStorage.getItem(KEY_TOKEN);

  const logged = useMemo(() => {
    return !!userInfo.token;
  }, [userInfo]);

  const isAdmin = useMemo(() => {
    return logged && userInfo.role?.includes(UserRole.ADMIN);
  }, [logged, userInfo]);

  const { data: newUser } = useQuery({
    queryKey: ['userInfo'],
    queryFn: () => getApiUserInfo(),
    enabled: logged && !!storageToken,
  });

  useEffect(() => {
    if (newUser) {
      updateUser({ ...newUser.data });
    }
  }, [newUser, updateUser]);

  const toLogin = (user: User, onOk?: () => void) => {
    setUserState(user);
    updateUser(user);
    localStorage.setItem(KEY_TOKEN, user.token || '');
    Cookies.set('token', user.token || '');
    onOk?.();
  };

  const logout = (onOk?: () => void) => {
    resetInfo();
    setUserState({} as User);
    localStorage.removeItem(KEY_TOKEN);
    Cookies.remove('token');
    onOk?.();
  };

  return (
    <AuthContext.Provider value={{ logout, isAdmin, user: userState, logged, toLogin }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext };

export default AuthProvider;
