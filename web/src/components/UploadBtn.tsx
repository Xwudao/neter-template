import UploadCard from '@/components/UploadCard.tsx';
import { IconImage } from '@douyinfe/semi-icons';
import { Button, Popover, Typography } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

type IUploadBtn = {
  prefix?: string;
};
const UploadBtn: FC<PropsWithChildren<IUploadBtn>> = ({ prefix = 'images' }) => {
  console.log('uploadBtn render...');
  const [show, setShow] = useState(false);
  const [result, setResult] = useState<{
    key: string;
    url: string;
  }>();
  return (
    <>
      <Popover
        showArrow
        position={`bottomRight`}
        trigger={`custom`}
        visible={show}
        content={
          <div className={`p2 w-60 max-w-60`}>
            <h2 className={`fw-bold text-base`}>上传图片</h2>
            <UploadCard
              initValues={{ prefix }}
              onSuccess={(rtn) => {
                setResult(rtn.data);
              }}
            />
            {result?.url && (
              <div className={`flex flex-col gap-2 mt3`}>
                <Typography.Text className={`whitespace-pre-wrap break-all`} copyable={{ content: result?.url }}>
                  {result?.url}
                </Typography.Text>
                <Typography.Text className={`whitespace-pre-wrap break-all`} copyable={{ content: result?.key }}>
                  {result?.key}
                </Typography.Text>
              </div>
            )}
          </div>
        }>
        <Button icon={<IconImage />} onClick={() => setShow(!show)} />
      </Popover>
    </>
  );
};

export default UploadBtn;
