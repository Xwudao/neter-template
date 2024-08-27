interface BaseType<T> {
  id: number;
  parent_id: number | null;
  children?: T[];
  [key: string]: any;
}
function arrayToTree<T extends BaseType<T>>(nodes: T[]): T[] {
  const map = new Map<number, T>();
  const rootNodes: T[] = [];

  // 构建映射
  nodes.forEach((node) => {
    map.set(node.id, node);
  });

  // 构建树
  nodes.forEach((node) => {
    if (node.parent_id === null || node.parent_id === undefined) {
      rootNodes.push(node);
    } else {
      const parent = map.get(node.parent_id);
      if (parent) {
        (parent.children || (parent.children = [])).push(node);
      }
    }
  });

  return rootNodes;
}
export { arrayToTree };
