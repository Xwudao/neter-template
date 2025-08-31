import KindTypes from '@/core/kind_types';
import FriendLinkForm from './FriendLinkForm';
import ValuesForm from './ValuesForm';
import BackupLinkForm from '@/components/others/BackupLinkForm';

interface Props {
  kind: KindTypes;
  value?: string;
  onChange?: (value: string) => void;
  onKeyChange?: (key: string) => void;
}

function DataKindsBuilder(props: Props) {
  const { kind, value, onChange, onKeyChange } = props;
  console.log('🚀 ~ DataKindsBuilder ~ value:', value);

  const initValues = value ? JSON.parse(value)['values'] : [];

  const renderValuesForm = () => (
    <ValuesForm
      field={'values'}
      onKeyChange={onKeyChange}
      handleFormChange={(values) => {
        onChange?.(JSON.stringify(values));
      }}
      initValue={initValues}
    />
  );

  const renderFormByKind = () => {
    switch (kind) {
      case KindTypes.FriendLink:
        return <FriendLinkForm value={value} onChange={onChange} onKeyChange={onKeyChange} />;
      case KindTypes.BackupLink:
        return <BackupLinkForm value={value} onChange={onChange} onKeyChange={onKeyChange} />;
      case KindTypes.BanUA:
        return renderValuesForm();
      default:
        return <div>暂不支持该类型的配置</div>;
    }
  };

  return <div className="data-kinds-builder">{renderFormByKind()}</div>;
}

export default DataKindsBuilder;
