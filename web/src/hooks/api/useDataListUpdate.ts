import { useMutation } from '@tanstack/react-query';
import { DataList, postApiUpdateDataList } from '@/api/dataListApi';
import { onError, onSuccess } from '@/core/callback';

const useDataListUpdate = () => {
  const { mutate: updateDataList, isPending: updating } = useMutation({
    mutationKey: ['update-list-store'],
    mutationFn: postApiUpdateDataList,
  });

  const updateProp = (
    datalist: DataList,
    prop: string,
    updater: (currentValue: any) => any,
    onRefresh?: () => void,
  ) => {
    const jsonData = JSON.parse(datalist.value);
    const newValue = { ...jsonData, [prop]: updater(jsonData[prop]) };

    updateDataList(
      {
        key: datalist.key,
        id: datalist.id,
        value: JSON.stringify(newValue),
      },
      {
        onSuccess: onSuccess(false, onRefresh),
        onError: onError(),
      },
    );
  };

  return { updateProp, updating };
};

export default useDataListUpdate;
