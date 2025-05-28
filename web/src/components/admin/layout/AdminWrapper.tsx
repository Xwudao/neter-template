import { FC, PropsWithChildren, ReactNode } from 'react';
import classes from './styles.module.scss';
import AdminContent from '@/components/admin/layout/AdminContent';
import AdminToolbar from '@/components/admin/layout/AdminToolbar';

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
