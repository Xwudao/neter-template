import { Category, postApiCreateCategory } from '@/api/categoryApi.ts';
import { onError, onSuccess } from '@/core/callback.ts';
import CategoryForm from '@/pages/admin/popovers/CategoryForm.tsx';
import FormWrapper from '@/pages/admin/popovers/FormWrapper.tsx';
import { Button } from '@douyinfe/semi-ui';
import { useMutation } from '@tanstack/react-query';
import { FC, PropsWithChildren } from 'react';

type IAddCategory = {
  categories: Category[];
  onAdded?: () => void;
};
const AddCategory: FC<PropsWithChildren<IAddCategory>> = ({ categories, onAdded }) => {
  console.log('addCategory render...');
  const { mutate: doAdd, isPending } = useMutation({
    mutationFn: postApiCreateCategory,
  });
  const handleAdd = (values: any) => {
    doAdd(values, {
      onSuccess: onSuccess(true, () => onAdded?.()),
      onError: onError(),
    });
  };
  return (
    <div className={`p2 w-62`}>
      <FormWrapper title={`添加分类`} onSubmit={handleAdd}>
        <CategoryForm categories={categories} />
        <Button block loading={isPending} htmlType={`submit`}>
          提交
        </Button>
      </FormWrapper>
    </div>
  );
};

export default AddCategory;
