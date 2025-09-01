import KindTypes, { BanUAValue } from '@/core/kind_types';
import { useQuery } from '@tanstack/react-query';
import clsx from 'clsx';
import classes from './data-list.module.scss';
import BoxLoading from '@/components/loading/BoxLoading';
import { getAdminApiGetDataListSortData, DataList } from '@/api/dataListApi';

interface Props {
  kind_type: KindTypes;
  className?: string;
  onToggleStatus?: (item: DataList, currentStatus: boolean) => void;
}

function DataListViewer(props: Props) {
  const { kind_type, className, onToggleStatus } = props;

  const { data, isLoading, error } = useQuery({
    queryKey: ['dataListSortData', kind_type],
    queryFn: () => getAdminApiGetDataListSortData({ kind: kind_type }),
    enabled: !!kind_type,
  });

  if (isLoading) {
    return <BoxLoading />;
  }

  if (error) {
    return <div className={classes.errorState}>加载失败: {error.message}</div>;
  }

  const dataList = data?.data || [];

  if (dataList.length === 0) {
    return (
      <div className={clsx(classes.container, className)}>
        <div className={classes.emptyState}>暂无数据</div>
      </div>
    );
  }

  const handleStatusClick = (item: DataList, currentStatus: boolean) => {
    if (onToggleStatus) {
      onToggleStatus(item, currentStatus);
    }
  };

  return (
    <div className={clsx(classes.container, className)}>
      <div className={classes.dataCount}>共 {dataList.length} 条数据</div>
      {dataList.map((item) => {
        if (kind_type === KindTypes.BanUA) {
          try {
            const banUAValue: BanUAValue = JSON.parse(item.value);
            return (
              <div key={item.id} className={classes.dataContainer}>
                <div className={classes.dataItem}>
                  <div className={classes.itemHeader}>
                    <span className={classes.itemId}>ID: {item.id}</span>
                    <div className={classes.statusContainer}>
                      <span
                        className={clsx(
                          classes.statusBadge,
                          classes.statusClickable,
                          banUAValue.enable ? classes.statusEnabled : classes.statusDisabled,
                        )}
                        onClick={() => handleStatusClick(item, banUAValue.enable)}>
                        {banUAValue.enable ? '启用' : '禁用'}
                      </span>
                    </div>
                  </div>
                  <div className={classes.valuesContainer}>
                    <div className={classes.valuesLabel}>内容:</div>
                    <div className={classes.valuesText}>{banUAValue.values.join('、')}</div>
                  </div>
                </div>
              </div>
            );
            // oxlint-disable-next-line no-unused-vars
          } catch (e) {
            return (
              <div key={item.id} className={classes.errorItem}>
                <span className={classes.errorLabel}>数据格式错误 (ID: {item.id}):</span>
                <div className={classes.errorValue}>{item.value}</div>
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
