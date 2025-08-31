import KindTypes from '@/core/kind_types';
import { IconEyeOpened } from '@douyinfe/semi-icons';
import { Button, Modal } from '@douyinfe/semi-ui';
import { useState } from 'react';
import clsx from 'clsx';
import classes from './DataViewBtn.module.scss';
import DataListViewer from '@/components/others/DataListViewer';

interface Props {
  kind_type: KindTypes;
  className?: string;
  disabled?: boolean;
}

function DataViewBtn(props: Props) {
  const { kind_type, className = '', disabled } = props;
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => setIsModalOpen(false);

  return (
    <>
      <Button
        onClick={openModal}
        disabled={disabled}
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
        width={320}
        bodyStyle={{ maxHeight: '50vh', overflow: 'auto' }}>
        <DataListViewer kind_type={kind_type} />
      </Modal>
    </>
  );
}

export default DataViewBtn;
