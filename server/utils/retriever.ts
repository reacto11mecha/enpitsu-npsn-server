export const getSpecificSchool = defineCachedFunction(
  async (npsn: npsn) => {
    const data = await useStorage("assets:server").getItem("schools.json");

    const parsedData = schoolsDataScheme.parse(data);

    const specificSchool = parsedData.find((school) => school.npsn === npsn);

    return specificSchool ?? null;
  },
  {
    maxAge: 60 * 60 * 2,
    name: "npsnKey",
    getKey: (npsn: number) => npsn,
  },
);
