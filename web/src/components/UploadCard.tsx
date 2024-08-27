import { UploadS3Res } from '@/api/uploadApi.ts';
import { KEY_TOKEN } from '@/core/constants.ts';
import { API_URL } from '@/core/types.ts';
import { generateRandomString } from '@/utils/string.ts';
import { Button, Form, Typography } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

const { Upload, Input } = Form;
type IUploadCard = {
  initValues?: {
    prefix?: string;
    object?: string;
  };
  onSuccess?: (rtn: UploadS3Res) => void;
};
const UploadCard: FC<PropsWithChildren<IUploadCard>> = ({ onSuccess, initValues }) => {
  console.log('uploadCard render...');
  // const [percent, setPercent] = useState(0);
  const action = useMemo(() => {
    return API_URL + '/admin/v1/upload/s3';
  }, []);

  const ref = useRef<any>(null!);

  const handleSubmit = () => {
    ref.current.upload();
  };

  return (
    <>
      <Form
        initValues={initValues}
        onSubmit={handleSubmit}
        render={({ formState, formApi }) => {
          return (
            <>
              <Input
                field={`prefix`}
                placeholder={`前缀`}
                label={`前缀(path)`}
                rules={[{ required: true, message: '必填' }]}
              />
              <Input
                field={`object`}
                placeholder={`文件名，不含后缀`}
                label={`文件名(object)`}
                rules={[{ required: true, message: '必填' }]}
                extraText={
                  <>
                    <Typography.Text
                      size={`small`}
                      link
                      className={`cp`}
                      onClick={() => {
                        formApi.setValue('object', generateRandomString(12));
                      }}>
                      随机生成
                    </Typography.Text>
                  </>
                }
              />
              <Upload
                rules={[{ required: true, message: '必填' }]}
                label={`上传文件(file)`}
                uploadTrigger="custom"
                field={`file`}
                ref={ref}
                action={action}
                data={{
                  prefix: formState.values['prefix'],
                  object: formState.values['object'],
                }}
                draggable={true}
                headers={{
                  Authorization: `Bearer ${localStorage.getItem(KEY_TOKEN)}`,
                }}
                onSuccess={(rtn: UploadS3Res) => {
                  onSuccess?.(rtn);
                }}
                addOnPasting={true}
                // onProgress={(percent) => setPercent(percent)}
                fileName={`file`}
                dragMainText={'点击上传文件或拖拽文件到这里'}
                dragSubText="支持任意类型文件"
              />

              {/*{percent > 0 && percent < 100 && (*/}
              {/*  <Progress percent={percent} aria-label="upload precent" className={`mb3`} />*/}
              {/*)}*/}

              <Button htmlType={`submit`} block>
                上传
              </Button>
            </>
          );
        }}
      />
    </>
  );
};

export default UploadCard;
