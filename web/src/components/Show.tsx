import { FC, PropsWithChildren } from 'react';

type IShow = {
  show?: boolean;
};
const Show: FC<PropsWithChildren<IShow>> = ({ children, show = true }) => {
  if (!show) return null;
  return <>{children}</>;
};

export default Show;
