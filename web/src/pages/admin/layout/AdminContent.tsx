import classes from './pages.module.scss';
import { FC, PropsWithChildren } from 'react';

type IAdminContent = {};
const AdminContent: FC<PropsWithChildren<IAdminContent>> = ({ children }) => {
  console.log('adminContent render...');
  return <div className={classes.adminContent}>{children}</div>;
};

export default AdminContent;
