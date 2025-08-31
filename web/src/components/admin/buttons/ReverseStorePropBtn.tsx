import { DataList } from '@/api/dataListApi';
import AppIcon from '@/components/AppIcon';
import useDataListUpdate from '@/hooks/api/useDataListUpdate';
import { Button, Tooltip } from '@douyinfe/semi-ui';

interface Props {
  datalist: DataList;
  size?: 'small' | 'default';
  prop: string;
  name: string;
  onRefresh?: () => void;
}

function ReverseStorePropBtn(props: Props) {
  const { datalist, prop, name, size = 'default', onRefresh } = props;

  const { updateProp, updating } = useDataListUpdate();

  return (
    <Tooltip content={`反转属性: ${name}`}>
      <Button
        loading={updating}
        onClick={() => {
          updateProp(datalist, prop, (currentValue) => !currentValue, onRefresh);
        }}
        size={size}
        icon={<AppIcon size={'small'} icon={'i-icon-park-outline-reverse-operation-out'} />}></Button>
    </Tooltip>
  );
}

export default ReverseStorePropBtn;
