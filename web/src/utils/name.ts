const extractTitleName = (data: any) => {
  try {
    const { title, name, keyword, word } = JSON.parse(data) || {};
    return title || name || keyword || word || null;
  } catch (error) {
    console.error('Error parsing JSON:', error);
    return null;
  }
};

const reverseProp = (data: any, prop: string) => {
  //if prop is boolean, and exists, return the opposite value
  try {
    const jsonData = JSON.parse(data);
    if (typeof jsonData[prop] === 'boolean') {
      jsonData[prop] = !jsonData[prop];
    }

    return JSON.stringify(jsonData);
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }

  return data; // Return original data if parsing fails
};

const hasProp = (data: any, prop: string) => {
  try {
    const jsonData = JSON.parse(data);
    return Object.prototype.hasOwnProperty.call(jsonData, prop);
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }
  return false;
};

const hasShow = (data: any, props: string[]) => {
  // if prop is boolean, and exists, return it's value, else undefined
  try {
    const jsonData = JSON.parse(data);
    for (const p of props) {
      if (typeof jsonData[p] === 'boolean') {
        return jsonData[p];
      }
    }
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }
  return undefined; // Return undefined if parsing fails or prop is not boolean
};

export { extractTitleName, reverseProp, hasProp, hasShow };
