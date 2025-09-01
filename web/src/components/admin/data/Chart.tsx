import { RadioGroup, Radio } from '@douyinfe/semi-ui';
import { VChart } from '@visactor/react-vchart';

const commonSpec = {
  type: 'bar',
  data: [
    {
      id: 'barData',
      values: [
        { type: 'Date', month: 'Monday', sales: 22 },
        { type: 'Drate', month: 'Tuesday', sales: 13 },
        { type: 'Date', month: 'Wednesday', sales: 25 },
        { type: 'Date', month: 'Thursday', sales: 29 },
        { type: 'Date', month: 'Friday', sales: 38 },
      ],
    },
  ],
  title: {
    visible: true,
    text: 'Bar chart',
    subtext: 'This is a bar chart',
  },
  legends: {
    visible: true,
  },
  xField: 'month',
  yField: 'sales',
  seriesField: 'type',
};

function Chart() {
  //   const {} = props;
  const [direction, setDirection] = useState('vertical');
  const onChange = useCallback((e: any) => setDirection(e.target.value), []);
  const spec = useMemo(() => {
    const isVertical = direction === 'vertical';
    return {
      ...commonSpec,
      xField: isVertical ? 'month' : 'sales',
      yField: isVertical ? 'sales' : 'month',
      direction: direction,
    };
  }, [direction]);
  return (
    <>
      <RadioGroup onChange={onChange} value={direction}>
        <Radio value={'vertical'}>vertical</Radio>
        <Radio value={'horizontal'}>horizontal</Radio>
      </RadioGroup>
      <div style={{ height: 440 }}>
        <VChart key={direction} spec={spec} />
      </div>
    </>
  );
}

export default Chart;
