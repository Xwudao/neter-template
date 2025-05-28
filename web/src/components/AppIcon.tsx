import Icon from '@douyinfe/semi-icons';
import { Tooltip } from '@douyinfe/semi-ui';
import type { IconProps } from '@douyinfe/semi-ui/lib/es/icons';
import clsx from 'clsx';
import { forwardRef } from 'react';

interface Props extends Omit<IconProps, 'svg' | 'icon' | 'ref' | 'spin'> {
  icon: string;
  className?: string;
  spin?: boolean;
  tooltip?: string;
  text?: string;
}

const AppIcon = forwardRef<HTMLElement, Props>((props, ref) => {
  const { icon, className, text, spin, tooltip, size = 'default', ...remaining } = props;

  const IconComponent = (
    <Icon
      ref={ref}
      {...(text?.trim() ? {} : remaining)}
      size={size}
      svg={
        <i
          ref={ref}
          className={clsx(icon, className, {
            'semi-icon-spinning': spin,
          })}
        ></i>
      }
    />
  );

  if (tooltip?.trim()) {
    return (
      <Tooltip content={tooltip} position={'top'}>
        {IconComponent}
      </Tooltip>
    );
  }

  if (text?.trim()) {
    return (
      <div className={clsx('inline-flex items-center', className)} {...remaining}>
        {IconComponent}
        <span>{text}</span>
      </div>
    );
  }

  return IconComponent;
});

AppIcon.displayName = 'AppIcon';

export default AppIcon;
