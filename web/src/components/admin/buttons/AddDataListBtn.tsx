import AddDataList from '@/components/admin/surface/AddDataList';
import { IconPlus } from '@douyinfe/semi-icons';
import { Button, Modal } from '@douyinfe/semi-ui';
import clsx from 'clsx';
import { useState } from 'react';

interface Props {
  className?: string;
  disabled?: boolean;
  size?: 'small' | 'default' | 'large';
  onSuccess?: () => void;
}

function AddDataListBtn(props: Props) {
  const { className = '', size = 'default', disabled, onSuccess } = props;
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => setIsModalOpen(false);

  const handleSuccess = () => {
    closeModal();
    onSuccess?.();
  };

  return (
    <>
      <Button
        onClick={openModal}
        disabled={disabled}
        type="primary"
        size={size}
        icon={<IconPlus />}
        className={clsx(className)}>
        添加数据
      </Button>

      <Modal
        title="添加数据"
        visible={isModalOpen}
        onCancel={closeModal}
        footer={<></>}
        width={320}
        bodyStyle={{ padding: 0 }}>
        <AddDataList onClose={closeModal} onSuccess={handleSuccess} />
      </Modal>
    </>
  );
}

export default AddDataListBtn;
