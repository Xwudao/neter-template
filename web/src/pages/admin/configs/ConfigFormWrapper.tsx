import { Form } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

type IConfigFormWrapper = {
  className?: string;
  initValues?: Record<string, any>;
  onSubmit: (values: Record<string, any>) => void;
};
const ConfigFormWrapper: FC<PropsWithChildren<IConfigFormWrapper>> = ({
  children,
  className = '',
  initValues,
  onSubmit,
}) => {
  console.log('configFormWrapper render...');
  return (
    <>
      <Form initValues={initValues} className={className} onSubmit={onSubmit}>
        {children}
      </Form>
    </>
  );
};

export default ConfigFormWrapper;
