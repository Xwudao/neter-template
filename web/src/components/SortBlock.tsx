import { EmptyFunc } from '@/core/types.ts';
import classnames from 'classnames';
import { FC, PropsWithChildren, useState, useEffect } from 'react';
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import {
  closestCenter,
  DndContext,
  DragEndEvent,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core';

type ISortBlock = {
  data: any[];
  labelKey: string;
  className?: string;
  onSort: (data: any[]) => void;
};
const SortBlock: FC<PropsWithChildren<ISortBlock>> = ({
  data,
  labelKey,
  className = '',
  onSort = EmptyFunc,
}) => {
  console.log('sortBlock render...');

  const [innerData, setInnerData] = useState([] as any[]);
  useEffect(() => {
    setInnerData([...(data || [])]);
  }, [data]);
  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    }),
  );
  function handleDragEnd(event: DragEndEvent) {
    const { active, over } = event;
    if (over) {
      if (active.id !== over.id) {
        setInnerData((items) => {
          const oldIndex = items.findIndex((item) => item.id === active.id);
          const newIndex = items.findIndex((item) => item.id === over.id);
          const ordered = arrayMove(items, oldIndex, newIndex);

          onSort(ordered);
          return ordered;
        });
      }
    }
  }
  return (
    <div className={classnames(className, 'flex flex-col items-start')}>
      <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
        <SortableContext items={innerData} strategy={verticalListSortingStrategy}>
          {innerData.map((item, idx) => (
            <SortItem key={item.id} data={item} labelKey={labelKey} />
          ))}
        </SortableContext>
      </DndContext>
    </div>
  );
};

type ISortItem = {
  data: any;
  labelKey: string;
};
const SortItem: FC<PropsWithChildren<ISortItem>> = ({ labelKey, data }) => {
  console.log('SortItem render...');
  const { attributes, listeners, setNodeRef, transform, transition } = useSortable({
    id: data.id,
  });
  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  return (
    <div
      ref={setNodeRef}
      className={`cmn-text cursor-grab select-none max-w-full inline-flex items-start`}
      style={style}
      {...attributes}
      {...listeners}>
      <i className="i-carbon-drag-vertical"></i>
      <span className={`am truncate flex-1 ml-1`}>{data[labelKey]}</span>
    </div>
  );
};

export default SortBlock;
