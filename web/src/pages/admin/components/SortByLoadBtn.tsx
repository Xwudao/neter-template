import Loading from '@/components/Loading.tsx';
import SortBlock from '@/components/SortBlock.tsx';
import { IconSort } from '@douyinfe/semi-icons';
import { Button, Divider, Popover } from '@douyinfe/semi-ui';
import { useQuery } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

export type DataType = {
  id: number;
  [key: string]: any;
};

type ISortByLoadBtn = {
  key: string;
  labelKey?: string;
  loading?: boolean;
  disabled?: boolean;
  loadData: () => Promise<{ data: DataType[] }>;

  onOk: (data: any[]) => void;
};
const SortByLoadBtn: FC<PropsWithChildren<ISortByLoadBtn>> = ({
  loadData,
  labelKey = 'name',
  disabled = false,
  loading = false,
  key,
  onOk,
}) => {
  console.log('sortByLoadBtn render...');

  const [show, setShow] = useState(false);

  const { data, isPending } = useQuery({
    queryKey: ['load-data', key],
    queryFn: () => loadData(),
    enabled: show,
  });

  const [inData, setInData] = useState(data?.data || []);
  useEffect(() => setInData(data?.data || []), [data]);

  return (
    <Popover
      showArrow
      visible={show}
      position={`bottomRight`}
      trigger={`custom`}
      content={
        <div className={`p2 min-w-40 space-y-3`}>
          <h3 className={`fw-bold text-base`}>更新顺序</h3>
          <Divider margin={10} />
          <Loading show={isPending} className={`w-full py-2 text-center`} />
          <SortBlock
            className={`p1 max-w-60 max-h-40 overflow-y-auto`}
            data={inData}
            labelKey={labelKey}
            onSort={(data) => setInData(data)}
          />
          <Button
            block
            disabled={isPending}
            loading={loading}
            onClick={() => {
              onOk(inData);
            }}>
            确定
          </Button>
        </div>
      }>
      <Button disabled={disabled} icon={<IconSort />} onClick={() => setShow(!show)} />
    </Popover>
  );
};

export default SortByLoadBtn;
