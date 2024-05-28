import { z } from "zod";

export const npsnScheme = z
  .union([z.string().min(8).max(8), z.number().int().positive()])
  .pipe(z.coerce.number().min(1));

export const schoolsDataScheme = z.array(
  z.object({
    npsn: npsnScheme,
    uri: z.string().url(),
  }),
);
