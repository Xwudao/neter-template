import classes from './styles.module.scss';
import { FC, PropsWithChildren } from 'react';

type IAdminContent = {};
const AdminContent: FC<PropsWithChildren<IAdminContent>> = ({ children }) => {
  return <div className={classes.adminContent}>{children}</div>;
};

export default AdminContent;
