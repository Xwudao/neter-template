import dayjs from 'dayjs';

const formatDate = (str: string) => {
  return dayjs(str).format('YYYY-MM-DD HH:mm:ss');
};

export { formatDate };
