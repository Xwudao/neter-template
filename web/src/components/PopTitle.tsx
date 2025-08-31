import { IconClose } from '@douyinfe/semi-icons';
import { Button, Divider } from '@douyinfe/semi-ui';
import clsx from 'clsx';

interface Props {
  title: string;
  onClose?: () => void;
  className?: string;
  divided?: boolean;
  titleSize?: 'small' | 'default' | 'large';
}

function PopTitle(props: Props) {
  const { title: text, onClose, titleSize = 'default', className, divided } = props;

  return (
    <div className={clsx(className)}>
      <div className={'flex justify-between items-center'}>
        <h3
          className={clsx('fw-bold', {
            'text-sm': titleSize === 'small',
            'text-lg': titleSize === 'large',
            'text-base': titleSize === 'default',
          })}
        >
          {text}
        </h3>
        {onClose && (
          <Button
            icon={<IconClose />}
            size={'small'}
            theme={'borderless'}
            type={'tertiary'}
            onClick={onClose}
          />
        )}
      </div>
      {divided && <Divider margin={10} />}
    </div>
  );
}

export default PopTitle;
