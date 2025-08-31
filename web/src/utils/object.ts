const hasEnable = (jsonValue: string): boolean | undefined => {
  try {
    const parsed = JSON.parse(jsonValue);
    if (!Object.prototype.hasOwnProperty.call(parsed, 'enable')) {
      return undefined; // No 'enable' key in the object
    }
    return parsed.enable === true; // Return true or false based on the value
  } catch (error) {
    console.error('Invalid JSON format:', error);
    return false; // Invalid JSON format
  }
};

export { hasEnable };
