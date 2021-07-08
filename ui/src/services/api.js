import { stringify } from 'qs';
import request from '@/utils/request';

// About Us API
export async function queryAboutUsList(params) {
  return request(`/api/v1/about-us?${stringify(params)}`);
}

export async function addAboutUs(params) {
  return request(`/api/v1/about-us?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateAboutUs(params) {
  return request(`/api/v1/about-us/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeAboutUs(params) {
  return request(`/api/v1/about-us/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Target Protein API
export async function queryTargetProteinList(params) {
  return request(`/api/v1/technology/target-protein?${stringify(params)}`);
}

export async function addTargetProtein(params) {
  return request(`/api/v1/technology/target-protein?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateTargetProtein(params) {
  return request(`/api/v1/technology/target-protein/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeTargetProtein(params) {
  return request(`/api/v1/technology/target-protein/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// CADD API
export async function queryCADDList(params) {
  return request(`/api/v1/technology/cadd?${stringify(params)}`);
}

export async function addCADD(params) {
  return request(`/api/v1/technology/cadd?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateCADD(params) {
  return request(`/api/v1/technology/cadd/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeCADD(params) {
  return request(`/api/v1/technology/cadd/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// SBDD API
export async function querySBDDList(params) {
  return request(`/api/v1/technology/sbdd?${stringify(params)}`);
}

export async function addSBDD(params) {
  return request(`/api/v1/technology/sbdd?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateSBDD(params) {
  return request(`/api/v1/technology/sbdd/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeSBDD(params) {
  return request(`/api/v1/technology/sbdd/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// DEL API
export async function queryDELList(params) {
  return request(`/api/v1/technology/del?${stringify(params)}`);
}

export async function addDEL(params) {
  return request(`/api/v1/technology/del?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateDEL(params) {
  return request(`/api/v1/technology/del/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeDEL(params) {
  return request(`/api/v1/technology/del/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Products API
export async function queryProductsList(params) {
  return request(`/api/v1/products?${stringify(params)}`);
}

export async function addProduct(params) {
  return request(`/api/v1/products?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateProduct(params) {
  return request(`/api/v1/products/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeProduct(params) {
  return request(`/api/v1/products/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Partnering API
export async function queryPartneringList(params) {
  return request(`/api/v1/partnering?${stringify(params)}`);
}

export async function addPartnering(params) {
  return request(`/api/v1/partnering?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updatePartnering(params) {
  return request(`/api/v1/partnering/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removePartnering(params) {
  return request(`/api/v1/partnering/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Pipeline API
export async function queryPipelineList(params) {
  return request(`/api/v1/pipeline?${stringify(params)}`);
}

export async function addPipeline(params) {
  return request(`/api/v1/pipeline?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updatePipeline(params) {
  return request(`/api/v1/pipeline/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removePipeline(params) {
  return request(`/api/v1/pipeline/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// News API
export async function queryNewsList(params) {
  return request(`/api/v1/news?${stringify(params)}`);
}

export async function addNews(params) {
  return request(`/api/v1/news?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateNews(params) {
  return request(`/api/v1/news/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeNews(params) {
  return request(`/api/v1/news/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Careers API
export async function queryCareersList(params) {
  return request(`/api/v1/careers?${stringify(params)}`);
}

export async function addCareers(params) {
  return request(`/api/v1/careers?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateCareers(params) {
  return request(`/api/v1/careers/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeCareers(params) {
  return request(`/api/v1/careers/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}


