import { Category } from '@/api/categoryApi.ts';
import { arrayToTree } from '@/utils/tree.ts';
import { Form } from '@douyinfe/semi-ui';
import { FC, PropsWithChildren } from 'react';

const { TreeSelect, TextArea, Input, Slider } = Form;

type ICategoryForm = {
  category?: Category;
  categories: Category[];
};
const CategoryForm: FC<PropsWithChildren<ICategoryForm>> = ({ category, categories }) => {
  console.log('categoryForm render...');

  const treeData = useMemo(() => {
    const data = categories.map((item) => ({
      parent_id: item.parent_id,
      label: item.name,
      value: item.id,
      id: item.id,
      key: `${item.parent_id || 0}-${item.id}`,
    }));
    return arrayToTree(data);
  }, [categories]);

  return (
    <>
      <Input field={`name`} label={`名称`} placeholder={`请输入名称`} rules={[{ required: true, message: '必填' }]} />
      <Input field={`image`} label={`图片`} placeholder={`请输入图片地址`} />
      <TextArea field={`description`} label={`描述`} placeholder={`请输入描述`} />

      <TreeSelect
        className={`w-full ovh`}
        placeholder={`请选择父分类`}
        field={`pid`}
        label={`父分类`}
        defaultExpandAll
        maxTagCount={1}
        checkRelation="unRelated"
        treeData={treeData}
      />

      <Slider label={`权重排序`} field={`item_order`} min={1} max={100} showBoundary={false} />
    </>
  );
};

export default CategoryForm;
