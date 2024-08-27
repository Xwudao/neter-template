import { User } from '@/api/userApi.ts';
import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

type UserState = {
  userInfo: User;
  updateUser: (userInfo: Partial<User>) => void;

  resetInfo: () => void;
};

export const initUserInfo = (): User => ({
  id: 0,
  role: '',
  create_time: '',
  update_time: '',
  username: '',
  token: '',
});

const useUserInfo = create<UserState>()(
  devtools(
    persist(
      (set, get) => ({
        userInfo: initUserInfo(),
        updateUser: (userInfo: Partial<User>) =>
          set({ userInfo: { ...get().userInfo, ...userInfo } }),
        resetInfo: () => set({ userInfo: initUserInfo() }),
      }),
      {
        name: 'user-info',
      },
    ),
  ),
);

export default useUserInfo;
