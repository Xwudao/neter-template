import { postApiWriteStaticFile } from '@/api/siteApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import { IconForward, IconLoading, IconSave } from '@douyinfe/semi-icons';
import { Form, Typography, useFormState } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import classnames from 'classnames';
import { FC, PropsWithChildren } from 'react';

const { TextArea } = Form;

type ISEOForm = { className?: string };
const SEOForm: FC<PropsWithChildren<ISEOForm>> = ({ className }) => {
  console.log('sEOForm render...');

  const { mutate: doWrite, isPending } = useMutation({
    mutationFn: postApiWriteStaticFile,
  });

  const handleWrite = (filename: string, data: string) => {
    doWrite(
      {
        filename: filename,
        data: data,
      },
      {
        onSuccess: onSuccess('保存成功'),
        onError: onError(),
      },
    );
  };

  const formState = useFormState();

  return (
    <div className={classnames(className)}>
      <TextArea
        field={`robots`}
        label={`Robots`}
        placeholder={`请输入Robots规则`}
        autosize={{ minRows: 1, maxRows: 10 }}
        rules={[{ required: true, message: '必填' }]}
        extraText={
          <p className={`flex gap-2`}>
            <Typography.Text
              size={`small`}
              link
              icon={!isPending ? <IconSave /> : <IconLoading spin />}
              onClick={() => handleWrite('robots.txt', formState.values['robots'])}>
              保存文件
            </Typography.Text>
            <Typography.Text size={`small`} link={{ href: '/robots.txt', target: '_blank' }} icon={<IconForward />}>
              访问
            </Typography.Text>
          </p>
        }
      />
    </div>
  );
};

export default SEOForm;
