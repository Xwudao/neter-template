import { Form } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

type IFormWrapper = {
  title: string;
  className?: string;
  initValues?: Record<string, any>;
  onSubmit: (values: Record<string, any>) => void;
};
const FormWrapper: FC<PropsWithChildren<IFormWrapper>> = ({
  children,
  className = '',
  title,
  initValues,
  onSubmit,
}) => {
  console.log('formWrapper render...', initValues);
  return (
    <div>
      <h5 className={`fw-bold text-base`}>{title}</h5>
      <Form initValues={initValues} className={className} onSubmit={onSubmit}>
        {children}
      </Form>
    </div>
  );
};

export default FormWrapper;
