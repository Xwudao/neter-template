import { FC, PropsWithChildren, useRef, useEffect, useState } from 'react';
import clsx from 'clsx';
import classes from './admin-sticky.module.scss';

type IAdminSticky = {
  /** 吸附位置，top: 顶部吸附，bottom: 底部吸附 */
  position?: 'top' | 'bottom';
  /** 距离顶部或底部的偏移量 */
  offset?: number;
  /** 层级控制 */
  zIndex?: number;
  /** 组件尺寸 */
  size?: 'small' | 'medium' | 'large';
  /** 自定义类名 */
  className?: string;
};

const AdminSticky: FC<PropsWithChildren<IAdminSticky>> = ({
  children,
  position = 'top',
  offset = 0,
  zIndex = 10,
  size = 'medium',
  className = '',
}) => {
  console.log('adminSticky render...');

  const [isStuck, setIsStuck] = useState(false);
  const stickyRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const element = stickyRef.current;
    if (!element) return;

    const observer = new IntersectionObserver(
      ([entry]) => {
        setIsStuck(!entry.isIntersecting);
      },
      {
        threshold: [1],
        rootMargin: position === 'top' ? `-${offset + 1}px 0px 0px 0px` : `0px 0px -${offset + 1}px 0px`,
      },
    );

    observer.observe(element);

    return () => observer.disconnect();
  }, [offset, position]);

  const stickyStyle = {
    [position]: `${offset}px`,
    zIndex,
  };

  return (
    <div
      ref={stickyRef}
      className={clsx(
        classes.adminSticky,
        classes[`sticky${position.charAt(0).toUpperCase() + position.slice(1)}`],
        classes[`size${size.charAt(0).toUpperCase() + size.slice(1)}`],
        { [classes.stuck]: isStuck },
        className,
      )}
      style={stickyStyle}>
      {children}
    </div>
  );
};

export default AdminSticky;
