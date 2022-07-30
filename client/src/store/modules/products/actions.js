export default {
  async loadProductPaginations(context) {
    const response = await fetch(
      `http://`+ process.env.VUE_APP_BASE_API_OSS+`/api/products?catalog_id[]=2&catalog_id[]=3&limit=20&page=1&catalog_id[]=4`
    );
    const responseData = await response.json();

    if (!response.ok) {
      const error = new Error(responseData.message || "Failed to fetch!");
      throw error;
    }

    context.commit("setProductPagination", responseData.data.rows);
    context.commit("setFetchTimestamp");
  },
};
