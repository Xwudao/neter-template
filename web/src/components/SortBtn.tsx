import SortBlock from '@/components/SortBlock.tsx';
import { IconSort } from '@douyinfe/semi-icons';
import { Button, Divider, Popover } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

type ISortBtn = {
  data: any[];
  labelKey: string;
  onOk: (data: any[]) => void;

  loading?: boolean;
  disabled?: boolean;
};
const SortBtn: FC<PropsWithChildren<ISortBtn>> = ({ data, disabled = false, loading, labelKey, onOk }) => {
  console.log('sortBtn render...', data);

  const [inData, setInData] = useState(data || []);
  useEffect(() => setInData(data), [data]);

  return (
    <Popover
      showArrow
      position={`bottomRight`}
      trigger={`click`}
      content={
        <div className={`p2 min-w-40 space-y-3`}>
          <h3 className={`fw-bold text-base`}>更新顺序</h3>
          <Divider margin={10} />
          <SortBlock
            className={`p1 max-w-60`}
            data={inData}
            labelKey={labelKey}
            onSort={(data) => setInData(data)}
          />
          <Button
            block
            loading={loading}
            onClick={() => {
              onOk(inData);
            }}>
            确定
          </Button>
        </div>
      }>
      <Button disabled={disabled} icon={<IconSort />} />
    </Popover>
  );
};

export default SortBtn;
