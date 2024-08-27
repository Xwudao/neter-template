import AdminContent from '@/pages/admin/layout/AdminContent.tsx';
import AdminToolbar from '@/pages/admin/layout/AdminToolbar.tsx';
import { FC, PropsWithChildren, ReactNode } from 'react';
import classes from './pages.module.scss';

type IAdminWrapper = {
  toolbar?: ReactNode;
};
const AdminWrapper: FC<PropsWithChildren<IAdminWrapper>> = ({ children, toolbar }) => {
  console.log('adminWrapper render...');
  return (
    <div className={classes.adminWrapper}>
      {toolbar && <AdminToolbar>{toolbar}</AdminToolbar>}
      <AdminContent>{children}</AdminContent>
    </div>
  );
};

export default AdminWrapper;
