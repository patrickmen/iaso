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

// Target Validation API
export async function queryTargetValidationList(params) {
  return request(`/api/v1/technology/target-validation?${stringify(params)}`);
}

export async function addTargetValidation(params) {
  return request(`/api/v1/technology/target-validation?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateTargetValidation(params) {
  return request(`/api/v1/technology/target-validation/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeTargetValidation(params) {
  return request(`/api/v1/technology/target-validation/${params.id}?lang=${params.lang}`, {
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

// Biomarker API
export async function queryBiomarkerList(params) {
  return request(`/api/v1/technology/biomarker-development?${stringify(params)}`);
}

export async function addBiomarker(params) {
  return request(`/api/v1/technology/biomarker-development?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateBiomarker(params) {
  return request(`/api/v1/technology/biomarker-development/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeBiomarker(params) {
  return request(`/api/v1/technology/biomarker-development/${params.id}?lang=${params.lang}`, {
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

// Academic Institution API
export async function queryAcademicInstitutionList(params) {
  return request(`/api/v1/partnering/academic-institution?${stringify(params)}`);
}

export async function addAcademicInstitution(params) {
  return request(`/api/v1/partnering/academic-institution?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateAcademicInstitution(params) {
  return request(`/api/v1/partnering/academic-institution/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeAcademicInstitution(params) {
  return request(`/api/v1/partnering/academic-institution/${params.id}?lang=${params.lang}`, {
    method: 'DELETE',
    body: {
      method: 'delete',
    },
  });
}

// Industrial Institution API
export async function queryIndustrialInstitutionList(params) {
  return request(`/api/v1/partnering/industrial-institution?${stringify(params)}`);
}

export async function addIndustrialInstitution(params) {
  return request(`/api/v1/partnering/industrial-institution?lang=${params.lang}`, {
    method: 'POST',
    body: {
      ...params,
      method: 'post',
    },
  });
}

export async function updateIndustrialInstitution(params) {
  return request(`/api/v1/partnering/industrial-institution/${params.id}?lang=${params.lang}`, {
    method: 'PUT',
    body: {
      ...params,
      method: 'update',
    },
  });
}

export async function removeIndustrialInstitution(params) {
  return request(`/api/v1/partnering/industrial-institution/${params.id}?lang=${params.lang}`, {
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


