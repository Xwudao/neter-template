import KindTypes, { BanUAValue } from '@/core/kind_types';
import { useQuery } from '@tanstack/react-query';
import clsx from 'clsx';
import classes from './DataListViewer.module.scss';
import { getAdminApiGetDataListSortData } from '@/api/dataListApi copy';
import BoxLoading from '@/components/loading/BoxLoading';

interface Props {
  kind_type: KindTypes;
  className?: string;
}

function DataListViewer(props: Props) {
  const { kind_type, className } = props;

  const { data, isLoading, error } = useQuery({
    queryKey: ['dataListSortData', kind_type],
    queryFn: () => getAdminApiGetDataListSortData({ kinds: [kind_type] }),
    enabled: !!kind_type,
  });

  if (isLoading) {
    return <BoxLoading />;
  }

  if (error) {
    return <div className={classes.errorState}>加载失败: {error.message}</div>;
  }

  const dataList = data?.data || [];

  return (
    <div className={clsx(classes.container, className)}>
      {dataList.map((item) => {
        if (kind_type === KindTypes.BanUA) {
          try {
            const banUAValue: BanUAValue = JSON.parse(item.value);
            return (
              <div key={item.id} className={classes.dataContainer}>
                <div className={classes.dataItem}>
                  <h3 className={classes.itemHeader}>所属ID: {item.id}</h3>
                  <div className={classes.statusContainer}>
                    <span className={classes.statusLabel}>状态: </span>
                    <span className={clsx(banUAValue.enable ? classes.statusEnabled : classes.statusDisabled)}>
                      {banUAValue.enable ? '启用' : '禁用'}
                    </span>
                  </div>
                  <div className={classes.valuesContainer}>
                    <span className={classes.valuesLabel}>内容:</span>
                    <div className={classes.valuesText}>{banUAValue.values.join('、')}</div>
                  </div>
                </div>
              </div>
            );
          } catch (e) {
            return (
              <div key={item.id} className={classes.errorItem}>
                数据格式错误: {item.value}
              </div>
            );
          }
        }
        return null;
      })}
    </div>
  );
}

export default DataListViewer;
