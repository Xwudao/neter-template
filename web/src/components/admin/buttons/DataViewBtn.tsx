import KindTypes from '@/core/kind_types';
import { IconEyeOpened } from '@douyinfe/semi-icons';
import { Button, Modal } from '@douyinfe/semi-ui';
import { useState } from 'react';
import clsx from 'clsx';
import classes from './data-view.module.scss';
import DataListViewer from '@/components/others/DataListViewer';
import useDataListUpdate from '@/hooks/api/useDataListUpdate';
import { DataList } from '@/api/dataListApi';
import { useQueryClient } from '@tanstack/react-query';

interface Props {
  kind_type: KindTypes;
  className?: string;
  disabled?: boolean;
  onRefresh?: () => void;
}

function DataViewBtn(props: Props) {
  const { kind_type, className = '', disabled, onRefresh } = props;
  const [isModalOpen, setIsModalOpen] = useState(false);
  const { updateProp, updating } = useDataListUpdate();
  const queryClient = useQueryClient();

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => {
    setIsModalOpen(false);
    onRefresh?.();
  };

  const handleToggleStatus = (item: DataList, currentStatus: boolean) => {
    updateProp(
      item,
      'enable',
      () => !currentStatus,
      () => {
        // Refresh the data after update
        queryClient.invalidateQueries({
          queryKey: ['dataListSortData', kind_type],
        });
      },
    );
  };

  return (
    <>
      <Button
        onClick={openModal}
        disabled={disabled || updating}
        type="primary"
        size={'small'}
        icon={<IconEyeOpened />}
        className={clsx(classes.dataViewBtn, className)}></Button>

      <Modal
        title="数据查看器"
        visible={isModalOpen}
        onCancel={closeModal}
        onOk={closeModal}
        okText="关闭"
        cancelButtonProps={{ style: { display: 'none' } }}
        width={340}
        bodyStyle={{
          maxHeight: '60vh',
          overflow: 'auto',
          padding: 0,
        }}>
        <DataListViewer kind_type={kind_type} onToggleStatus={handleToggleStatus} />
      </Modal>
    </>
  );
}

export default DataViewBtn;
